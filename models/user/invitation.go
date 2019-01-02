package user

import (
  "github.com/jinzhu/gorm"
)


//Invitation is the model used when sending invitations to a user.
type Invitation struct {
  gorm.Model
  Email string `gorm:"size:250;unique"`
  Token string `gorm:"size:32;unique"`
}

//NewInvitation accepts an email address and creates a new invitation.
func NewInvitation(email string) Invitation {
  invitation := Invitation{Email: email}

  return invitation
}
