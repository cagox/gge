package user

import (
  "fmt"
  "time"
  "encoding/gob"


  "github.com/jinzhu/gorm"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/crypto"

)

//User is the user model
type User struct {
  gorm.Model
  Email        string `gorm:"size:250;unique"` // Email address, also the users login name
  Password     string                          // Obvious. This will be a hash.
  IsAdmin      bool                            // Is the user an admin.

  isVerified   bool                            // Are they verified? There will be methods to set and test this.
  DateVerified time.Time                       // This timestamp will be set when the user is verified.

  Profile      Profile
}

//Profile is the model that will hold profile data
type Profile struct {
  gorm.Model

  UserID uint
  Name         string `gorm:"size:40"`         // The users Display name.
}

//UserForm is a struct to collect user data for validation.
type UserForm struct {
  Email    string
  Name     string
  Password string
}

func init() {
  gob.Register(User{})
  gob.Register(Profile{})
  gob.Register(UserForm{})
}

//GetUsers returns a list of all the users in the database.
func GetUsers() []User {
  var users []User
  if err := config.Config.Database.Find(&users).Error; err != nil {
    fmt.Println("Could Not Search for Users")
  }

  return users
}

//GetUserByEmail returns a user with the matching email.
func GetUserByEmail(email string) *User {
  user := &User{}
  config.Config.Database.Where("email = ?", email).First(&user)

  return user
}

//SetPassword sets the password on the user object
func (user *User) SetPassword(password string) {
  user.Password = crypto.HashPassword(password)
}

//Authenticate allows the login method to make sure we have the right person.
func (user *User) Authenticate(password string) bool {
  return crypto.ComparePassword(password, user.Password)
}

//ValidatePassword validates the offered user password.
func ValidatePassword(password string) []string {
  errors := make([]string, 0)
  if (len(password) < config.Config.MinPasswordLength) {
    errors = append(errors, "The password must be at least "+ string(config.Config.MinPasswordLength)+" characters.")
  }
  return errors
}

/*
ValidateUserForm validates the data provided and returns a slice of strings if there are any errors.
Its first argument is a filled in UserForm struct. This is the data to be validated/
Its second argument is a bool value declaring if the form is for a new user rather than just updating a user.
*/
func ValidateUserForm(newUser UserForm, isNew bool) []string {
  errors := make([]string, 0)
  var users []User
  //Validate Password:
  passwordErrors := ValidatePassword(newUser.Password)
  if passwordErrors != nil {
    for err := range passwordErrors {
      errors = append(errors, passwordErrors[err])
    }
  }

  //Validate Email Address:
  if isNew {
    config.Config.Database.Where("email = ?", newUser.Email).Find(&users)
    if (len(users) != 0) {
      errors = append(errors, "The email address "+newUser.Email+" already exists.")
    }
  }

  //Validate Display name
  if len(newUser.Name) < config.Config.MinimumNameLength {
    errors = append(errors, "The Display Name must be at least "+string(config.Config.MinimumNameLength)+"characters long.")
  }

  return errors
}

//CreateUserFromForm creates a new user object from the data provided via a UserForm object.
func CreateUserFromForm(newUser UserForm) (*User, *Profile) {
  profile := &Profile{Name: newUser.Name}
  user := &User{Email: newUser.Email, Password: crypto.HashPassword(newUser.Password)}
  return user, profile
}
