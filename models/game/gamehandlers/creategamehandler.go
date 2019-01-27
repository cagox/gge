package gamehandlers


import (
  "net/http"
  "html/template"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/models/game"
)


func createGameHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w,r)
  sessionData := ggsession.GetSessionData(session)
  pageData := ggsession.BasePageData{}
  pageData.BasicData(sessionData)
  pageData.Page = "CreateGame"



  if r.Method == "POST" {
    r.ParseForm()

    name := r.FormValue("name")
    description := r.FormValue("description")
    owner := sessionData.Email

    newGame := game.CreateGame(name, owner, description)

    err := game.InsertGame(newGame)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    sessionData.AddFlash("success", "Game Created")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  //Either Method != POST or there was an error creating the game.
  //Load the template.


  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/game/creategame.html")
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
