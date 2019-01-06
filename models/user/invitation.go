package user

import (
  "fmt"
  "math/rand"

  "github.com/globalsign/mgo/bson"

  "github.com/cagox/gge/config"
)


//Invitation is the model used when sending invitations to a user.
type Invitation struct {
  ID      bson.ObjectId `bson:"_id,omitempty"`
  Email   string
  Token   string
}

//NewInvitation accepts an email address and creates a new invitation.
func NewInvitation(email string) Invitation {
  invitation := Invitation{Email: email}
  return invitation
}

//InviteEmailExists returns true if an invite is already in the system.
func InviteEmailExists(email string) bool {
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  invitations := mongoSession.DB("gge").C("invitations")

  count, err := invitations.Find(bson.M{"email": email}).Count()

  if err != nil {
    //TODO: Add error handling.
  }

  if count > 0 {
    return true
  }
  return false
}


//InviteTokenExists returns true if an invite is already in the system.
func InviteTokenExists(token string) bool {
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  invitations := mongoSession.DB("gge").C("invitations")

  count, err := invitations.Find(bson.M{"token": token}).Count()

  if err != nil {
    //TODO: Add error handling.
  }

  if count > 0 {
    return true
  }
  return false
}

//InsertInvitation adds the provided invitation to the database.
func InsertInvitation(invite Invitation) error {
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  invitations := mongoSession.DB("gge").C("invitations")
  err := invitations.Insert(&invite)
  if err != nil {
    return err
  }
  return nil
}



const (
  characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

/*
GenerateInviteToken creates a token, n characters long, made up of the charactesr
listed in the characters string above.
*/
func GenerateInviteToken(n int) string {
  b := make([]byte, n)
  rand.Read(b)
  return fmt.Sprintf("%x", b)
}
