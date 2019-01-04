package user

import (
  //"github.com/jinzhu/gorm"
  "github.com/globalsign/mgo/bson"
)


//Invitation is the model used when sending invitations to a user.
type Invitation struct {
  ID           bson.ObjectId `bson:"_id,omitempty"`
  Email string
  Token string
}

//NewInvitation accepts an email address and creates a new invitation.
func NewInvitation(email string) Invitation {
  invitation := Invitation{Email: email}

  return invitation
}
