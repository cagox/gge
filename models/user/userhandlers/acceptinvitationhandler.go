package userhandlers

import (
  "net/http"
  "html/template"

  "github.com/gorilla/mux"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"

)

type acceptInvitePageData struct {
  ggsession.BasePageData
  Token string
}


func acceptInvitationHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := acceptInvitePageData{}
  pageData.Page = "AcceptInvitation"
  pageData.BasicData(sessionData)

  vars := mux.Vars(r)
  pageData.Token = vars["token"]


  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/user/acceptinvite.html")
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
