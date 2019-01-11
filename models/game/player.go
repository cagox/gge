package game


import (
  "github.com/globalsign/mgo/bson"
)

//Player acts as a subscription to a given game.
type Player struct {
  ID          bson.ObjectId `bson:"_id,omitempty"`
  GameID      bson.ObjectId //The game that this player is connected to.
  UserID      bson.ObjectId //The User that this player is connected to.
  IsPlayer    bool  // Is this Player a member of the game, or just a lurker?
  IsWatching  bool  // Are they actually watching the game?
  
}
