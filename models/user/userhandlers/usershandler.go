package userhandlers

import (
  "net/http"
  "html/template"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/models/user"
)

//UsersListData will hold the page data for the user list.
type usersListData struct {
  ggsession.BasePageData
  Profiles   []user.Profile
  PageNum    int
  IsNext     bool
  IsPrevious bool
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := usersListData{}
  pageData.Page = "Users"
  pageData.BasicData(sessionData)

  if ! pageData.Authenticated {
    sessionData.AddFlash("error", "You must be logged in to access the user list.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  if ! pageData.IsAdmin {
    sessionData.AddFlash("error", "Only Admins Should Look at the User List (At least for now).")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  pageData.Flashes = sessionData.GetFlashes(true)

  config.Config.Database.Find(&pageData.Profiles)

  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/user/users.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  t.Execute(w, pageData)
  return

}
