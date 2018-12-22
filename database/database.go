package database


import (
  //Database Stuff
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite" //Imports the sqlite3 driver.

  "github.com/cagox/gge/config"


  //Models
  "github.com/cagox/gge/models/user"

)


//OpenDatabase opens the database.
func OpenDatabase(runMigrations bool) (*gorm.DB) {
  //TODO: Change this to a case switch.
  //TODO: Figure out the RIGHT way to decide what to pass as aruments and how to pass arguments to Open.

  if config.Config.DatabaseType == "sqlite3" {
    database := openSQLiteDB(config.Config.DatabasePath)

    if runMigrations {RunMigrations(database)}

    return database
  }

  return nil
}


func openSQLiteDB(databasePath string) (*gorm.DB) {
  db, err := gorm.Open("sqlite3", databasePath)
  if err != nil {
    panic("Database could not be opened.")
  }

  return db
}

//RunMigrations runs gorms automated migration code.
func RunMigrations(database *gorm.DB) {
  //TODO: As I add these make sure that the migrations happen in a logical order. Example: Profile depends on User, so user has to go first.
  database.AutoMigrate(&user.User{})
  database.AutoMigrate(&user.Profile{})
}
