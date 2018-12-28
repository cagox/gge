package routes

import (
  //"fmt"
  "net/http"
  //"html/template"

  "github.com/cagox/gge/config"
  //"github.com/cagox/gge/session"
  "github.com/cagox/gge/models/user"

)

//Routes sets up the routes for the main package and then calls similar methods in the attached packages.
func Routes() {
  //Setup the main routes.
  setupMainRoutes()
  //Setup the routes for the connected packages.
  user.Routes()

}


func setupMainRoutes() {
  http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir(config.Config.StaticPath))))
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/logout", logoutHandler)
  http.HandleFunc("/admin/firstuser", firstUserHandle)
}
