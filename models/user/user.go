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
  ID           bson.ObjectId `bson:"_id,omitempty"`
  Email        string
  Password     string
  IsAdmin      bool
  Timestamp    time.Time
  LastUpdated  time.Time
  Profile      Profile
}

//Profile is the model that will hold profile data
type Profile struct {
  Name         string

  //Settings
  ItemsPerPage int
}



/*
type User struct {
  gorm.Model
  Email        string `gorm:"size:250;unique"` // Email address, also the users login name
  Password     string                          // Obvious. This will be a hash.
  IsAdmin      bool                            // Is the user an admin.

  isVerified   bool                            // Are they verified? There will be methods to set and test this.
  DateVerified time.Time                       // This timestamp will be set when the user is verified.

  Profile      Profile
}
*/

/*
type Profile struct {
  gorm.Model

  UserID       uint
  Name         string `gorm:"size:40"`         // The users Display name.

  //Settings
  ItemsPerPage int
}
*/

//Form is a struct to collect user data for validation.
type Form struct {
  Email    string
  Name     string
  Password string
}

//SafeUser is a version of user that is safe to send to the page.
type SafeUser struct {
  UserID  uint
  IsAdmin bool
}

func init() {
  gob.Register(User{})
  gob.Register(Profile{})
  gob.Register(Form{})
  gob.Register(Invitation{})
}



/*
DisplayName returns the Profile object's display name in the format of
Name#UserID where UserID is the uint associated with the object's UserID.
*/
/*
func (user *User) DisplayName() string {

  return user.Profile.Name+"#"+strconv.Itoa(int(user.ID))

}
*/
