package gamehandlers


import (
  "net/http"
  "html/template"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/models/game"

)

type gamesListData struct {
  ggsession.BasePageData
  Games      []game.Game
  PageNum    int
  IsNext     bool
  IsPrevious bool
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := gamesListData{}
  pageData.Page = "Games"
  pageData.BasicData(sessionData)


  //config.Config.Database.Find(&pageData.Profiles)
  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  collection :=  mongoSession.DB("gge").C("games")


  err := collection.Find(nil).Sort("game.name").All(&pageData.Games)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }



  t := template.New("base.html")
  t, err = t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/game/games.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  pageData.Flashes = sessionData.GetFlashes(true)
  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  t.Execute(w, pageData)
  return

}
