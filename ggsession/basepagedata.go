package ggsession


import (

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/models/user"



)


func init() {

}

//BasePageData is the data that most pages will need. This can be used to build the data struct for templates.
type BasePageData struct {
  Page          string
  Flashes       []Flash
  Authenticated bool
  IsAdmin       bool
}

//BasicData fills in the BasePageData struct from the provided session.
func (data *BasePageData)BasicData(session SessionData) {
  data.Authenticated = session.Authenticated

  user := user.User{}
  config.Config.Database.Where("id = ?", session.UserID).First(&user)
  data.IsAdmin = user.IsAdmin


  
}
