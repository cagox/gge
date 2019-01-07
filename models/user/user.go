package user

import (
  //"fmt"
  "time"
  "encoding/gob"
  //"strconv"
  //"net/http"

  //"github.com/jinzhu/gorm"

  "github.com/globalsign/mgo/bson"
)

//User is the user model
type User struct {
  ID              bson.ObjectId `bson:"_id,omitempty"`
  Email           string
  Password        string
  IsAdmin         bool
  EmailIsVerified bool
  Timestamp       time.Time
  LastUpdated     time.Time
  Profile         Profile
}

//Profile is the model that will hold profile data
type Profile struct {
  Name         string

  //Settings
  ItemsPerPage int
}

//Form is a struct to collect user data for validation.
type Form struct {
  Email    string
  Name     string
  Password string
}


func init() {
  gob.Register(User{})
  gob.Register(Profile{})
  gob.Register(Form{})
  gob.Register(Invitation{})
}
