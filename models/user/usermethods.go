package user

import (
  "fmt"
  //"net/http"

  //"github.com/jinzhu/gorm"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/crypto"
  //"github.com/cagox/gge/config"
  "github.com/globalsign/mgo/bson"
)



//GetUsers returns a list of all the users in the database.
func GetUsers() []User {
  var users []User
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()

  collection := mongoSession.DB("gge").C("users")

  err := collection.Find(nil).Sort("-timestamp").All(&users)
  if err != nil {
    fmt.Println(err) //TODO: Add proper error handling.
  }

  return users
}

/*
func GetUsers() []User {
  var users []User
  if err := config.Config.Database.Find(&users).Error; err != nil {
    fmt.Println("Could Not Search for Users")
  }

  return users
}
*/

//AreThereAnyUsers checks to see if the database has any users or not.
func AreThereAnyUsers() bool {
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  users := mongoSession.DB("gge").C("users")
  count, err := users.Count()
  if err != nil {
    //What the fuck?
    //TODO Add error handling
  }
  if count == 0 {
    return false
  }
  return true
}

//GetUserByEmail returns a user with the matching email.
/*
func GetUserByEmail(email string) *User {
  user := &User{}
  config.Config.Database.Where("email = ?", email).First(&user)

  return user
}
*/
func GetUserByEmail(email string) *User {
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  users := mongoSession.DB("gge").C("users")

  var user *User
  err := users.Find(bson.M{"email": email}).One(&user)
  if err != nil {
    //TODO Error handling
  }
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
func CreateUserFromForm(newUser Form) (*User) {
  profile := Profile{Name: newUser.Name, ItemsPerPage: 20}
  user := &User{Email: newUser.Email, Password: crypto.HashPassword(newUser.Password), Profile: profile}

  //profile.ItemsPerPage = 20

  return user
}
