package main

import (
  "fmt"
  "log"
  "net/http"

  //Database Stuff
  "github.com/cagox/gge/database"

  //"os"

  //Config related imports.
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/routes"
  //"github.com/cagox/gge/email"
  //"github.com/cagox/gge/models/user"

  )


func main() {
  database.DialMongoSession()
  defer config.Config.MongoSession.Close()

  //email.SystemEmail("burlingk@cagox.net", "Testing my new Function", "\n\nHopefuly it works!\n\n--Ken")

  routes.Routes()

  fmt.Println("A few vital settings:")
  fmt.Println("Static Path: "+config.Config.StaticPath)
  fmt.Println("Template Root: "+config.Config.TemplateRoot)
  fmt.Println("System Email: "+config.Config.FromName+" <"+config.Config.FromAddress+">")
  
  fmt.Println("")

  log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
