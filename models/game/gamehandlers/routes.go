package gamehandlers

import (
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
)

//Routes sets up routs for package user
func Routes() {
  config.Config.Router.HandleFunc("/games/create", ggsession.MustBeAuthenticated(createGameHandler))
  config.Config.Router.HandleFunc("/games", ggsession.MustBeAuthenticated(gamesHandler))
  config.Config.Router.HandleFunc("/game/{slug}", ggsession.MustBeAuthenticated(gameHandler))

}
