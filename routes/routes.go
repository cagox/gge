package routes

import (
  //"fmt"
  "net/http"
  //"html/template"

  "github.com/cagox/gge/config"
  //"github.com/cagox/gge/session"
  "github.com/cagox/gge/models/user/userhandlers"

)

//Routes sets up the routes for the main package and then calls similar methods in the attached packages.
func Routes() {
  //Setup the main routes.
  setupMainRoutes()
  //Setup the routes for the connected packages.
  userhandlers.Routes()

  http.Handle("/", config.Config.Router)

}


func setupMainRoutes() {
  http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir(config.Config.StaticPath))))
  config.Config.Router.HandleFunc("/", indexHandler)
  config.Config.Router.HandleFunc("/login", loginHandler)
  config.Config.Router.HandleFunc("/logout", logoutHandler)
  config.Config.Router.HandleFunc("/admin/firstuser", firstUserHandle)
}
