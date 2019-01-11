package game


import (
  "time"
  "strconv"
  "fmt"

  "github.com/globalsign/mgo/bson"
  "github.com/tv42/slug"

  "github.com/cagox/gge/config"
)


//Game is the model that other game related models hang off of.
type Game struct {
  ID          bson.ObjectId `bson:"_id,omitempty"`
  Name        string     //The Name of the Game
  Slug        string     //The slug used to identify the game in URLs
  Owner       string    //The email address of the Game's "owner"
  Description string    //The description of the game provided by its creator
  Timestamp   time.Time //The time that the game was created
}

//CreateGame takes the requested parameters and creates a new game instance.
func CreateGame(name string, owner string, description string) *Game {
  game := Game{ Name: name, Owner: owner, Description: description}
  game.Slug = generateGameSlug(game.Name)
  game.Timestamp = time.Now()

  return &game
}

func generateGameSlug(name string) string {
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  games := mongoSession.DB("gge").C("games")

  baseSlug := slug.Slug(name)
  done := false

  for i := 0; done == false; i++ {
    suffix := ""
    if i > 0 {
      suffix = strconv.Itoa(i)
    }
    newSlug := baseSlug+suffix
    test, err := games.Find(bson.M{"slug": newSlug}).Count()
    if err != nil {
      fmt.Println(err) //TODO: Proper Error HAndling.
    }

    if test == 0 {
      return newSlug
    }
  }








  return ""
}
