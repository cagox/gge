package userhandlers

import (
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
)

//Routes sets up routs for package user
func Routes() {
  config.Config.Router.HandleFunc("/profile", ggsession.MustBeAuthenticated(profileHandler))
  config.Config.Router.HandleFunc("/users", ggsession.MustBeAdmin(usersHandler))
  config.Config.Router.HandleFunc("/inviteuser", ggsession.MustBeAdmin(inviteUserHandler))
  config.Config.Router.HandleFunc("/users/invite/{token}", acceptInvitationHandler)
}
