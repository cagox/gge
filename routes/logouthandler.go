package routes

import (
  "net/http"
  //"fmt"

  "github.com/cagox/gge/ggsession"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w,r)
  sessionData := ggsession.GetSessionData(session)

  session.Values["sessiondata"] = ggsession.NewSessionData()
  sessionData.AddFlash("success", "Successfully logged out.")
  sessionData.Authenticated = false
  session.Values["sessiondata"] = sessionData

  session.Save(r, w)
  http.Redirect(w, r, "/", http.StatusSeeOther)
}
