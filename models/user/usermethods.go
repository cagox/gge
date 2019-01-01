package user

import (
  "fmt"
  //"net/http"

  //"github.com/jinzhu/gorm"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/crypto"
  //"github.com/cagox/gge/config"
)



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



//CreateUserFromForm creates a new user object from the data provided via a UserForm object.
func CreateUserFromForm(newUser Form) (*User, *Profile) {
  profile := &Profile{Name: newUser.Name}
  user := &User{Email: newUser.Email, Password: crypto.HashPassword(newUser.Password)}

  profile.ItemsPerPage = 20

  return user, profile
}
