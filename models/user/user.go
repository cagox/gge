package user

import (
  "time"

  "github.com/jinzhu/gorm"
  "encoding/gob"
)

//User is the user model
type User struct {
  gorm.Model
  Email        string `gorm:"size:250;unique"` // Email address, also the users login name
  Name         string `gorm:"size:40"`         // The users username. This will be the default displayname.
  password     string                          // Obvious. This will be a hash.
  isVerified   bool                            // Are they verified? There will be methods to set and test this.
  DateVerified time.Time                       // This timestamp will be set when the user is verified.
  IsAdmin      bool                            // Is the user an admin.

  Profile      Profile
}

//Profile is the model that will hold profile data
type Profile struct {
  gorm.Model

  UserID uint
}


func init() {
  gob.Register(User{})
  gob.Register(Profile{})
}
