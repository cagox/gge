package main

import (
  "fmt"
  "log"
  "net/http"

  //Database Stuff
  _ "github.com/jinzhu/gorm/dialects/sqlite" //Imports the sqlite3 driver.
  "github.com/cagox/gge/database"

  //"os"

  //Config related imports.
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/routes"
  //_ "github.com/cagox/gge/session"
  )


func main() {
  config.Config.Database = database.OpenDatabase(true)
  defer config.Config.Database.Close()

  routes.SetupRoutes()

  fmt.Println("\nDatabase Path: "+config.Config.DatabasePath)
  fmt.Println("Static Path: "+config.Config.StaticPath)
  fmt.Println("Template Root: "+config.Config.TemplateRoot)
  fmt.Println("")

  log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
