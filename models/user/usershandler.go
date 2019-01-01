package user

import (
  "net/http"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
)

//UsersListData will hold the page data for the user list.
type usersListData struct {
  ggsession.BasePageData
  Profiles   []Profile
  PageNum    int
  IsNext     bool
  IsPrevious bool
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := usersListData{}
  pageData.Page = "Users"
  pageData.Authenticated = sessionData.Authenticated

  if ! pageData.Authenticated {
    sessionData.AddFlash("error", "You must be logged in to access the user list.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  pageData.Flashes = sessionData.GetFlashes(true)

  config.Config.Database.Find(&pageData.Profiles)





}
