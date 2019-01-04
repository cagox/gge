package routes

import (
  //"fmt"
  "net/http"
  "html/template"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/models/user"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  //Test for empty database.
  //If database is Empty, go to Admin User Creation Screen.
  if (! user.AreThereAnyUsers()){
    http.Redirect(w, r, "/admin/firstuser", http.StatusSeeOther)
    return
  }

  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := ggsession.BasePageData{Page: "Index"}
  pageData.BasicData(sessionData)
  pageData.Flashes = sessionData.GetFlashes(true)

  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/base/index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  t.Execute(w, pageData)
  return
}
