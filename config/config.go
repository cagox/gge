//Package config handles the configuraion for the God Game Engine Program.
package config



import (
  "os"
  "github.com/mitchellh/go-homedir"
  "encoding/json"
  "github.com/jinzhu/gorm"
  //"github.com/gorilla/sessions"
)

//Config holds the system configuration.
var Config *GodGameConfiguration


//GodGameConfiguration holds the configuration information for the GGE program.
type GodGameConfiguration struct {
  DatabasePath      string
  StaticPath        string
  DatabaseType      string
  TemplateRoot      string
  AdminToken        string
  MinimumNameLength int
  MinPasswordLength int
  EncKey            string
  AuthKey           string
  SMTPServer        string
  FromAddress       string
  FromName          string
  SMTPPassword      string
  SMTPUserName      string

  Database      *gorm.DB  //Not in the Config File.
}

func init() {
  LoadConfiguration()
}



//LoadConfiguration Loads the configuration and sets the global variable.
func LoadConfiguration() {
  //TODO: Add error handling.
  Config = GetConfigs()
}

//GetConfigs returns a configuration struct.
func GetConfigs() *GodGameConfiguration {
  configpath := getConfigPath()
  ensureDirectory(configpath)

  configuration := GodGameConfiguration{}
  //err := gonfig.GetConf(configpath+"/gge.json", &configuration)

  file, err := os.Open(configpath+"/gge.json")

  //If there is an error (if the file doesn't exist or is malformed) we will created a default  configuration.
  if err != nil {
    //return defaultConfiguration(configpath)
    panic("Configuration File Did Not Open")
  }

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&configuration)

  if err != nil {
    //return defaultConfiguration(configpath)
    panic("Configuration File Malformed.")
  }

  return &configuration
}

//ensureDirectory() makes sure the given directory exists or return an error.
func ensureDirectory(base string) error {
  _, statErr := os.Stat(base)

  if statErr != nil {
    createErr := os.MkdirAll(base, 0755)
    if createErr != nil {
      panic("Could not ensure the configuration path.")
    }
  }

  return nil
}

//getConfigPath() Gets the config path from the environmental variable or uses a default.
func getConfigPath() string {
  //First we figure out where the configuration files should be.
  configpath, isEnv := os.LookupEnv("GGEROOT")
  if !isEnv {
    home, _ := homedir.Dir() //TODO: Add error checking.
    configpath = home+"/.config/gge"
  }
  return configpath
}
