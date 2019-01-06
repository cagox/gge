package userhandlers

import (
  "github.com/cagox/gge/config"
)

//Routes sets up routs for package user
func Routes() {
  config.Config.Router.HandleFunc("/profile", profileHandler)
  config.Config.Router.HandleFunc("/users", usersHandler)
  config.Config.Router.HandleFunc("/inviteuser", inviteUserHandler)
  config.Config.Router.HandleFunc("/users/invite/{token}", acceptInvitationHandler)
}
