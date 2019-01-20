package gamehandlers


import (
  "net/http"
)


func createGameHandler(w http.ResponseWriter, r *http.Request) {

  if r.Method == "POST" {
    //Attempt to create game.
  }

  //Either Method != POST or there was an error creating the game.
  //Load the template.



}
