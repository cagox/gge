package main

import (
  "fmt"
  "log"
  "net/http"

  //Database Stuff
  //_ "github.com/jinzhu/gorm/dialects/sqlite" //Imports the sqlite3 driver.
  "github.com/cagox/gge/database"

  //"os"

  //Config related imports.
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/routes"
  //"github.com/cagox/gge/email"

  )


func main() {
  database.DialMongoSession()
  defer config.Config.MongoSession.Close()

  //email.SystemEmail("burlingk@cagox.net", "Testing my new Function", "\n\nHopefuly it works!\n\n--Ken")

  routes.Routes()

  fmt.Println("Static Path: "+config.Config.StaticPath)
  fmt.Println("Template Root: "+config.Config.TemplateRoot)
  fmt.Println("")

  log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
