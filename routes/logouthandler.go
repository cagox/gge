package routes

import (
  "net/http"
  "fmt"

  "github.com/cagox/gge/ggsession"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w,r)
  sessionData := ggsession.GetSessionData(session)
  fmt.Println("Logout Handler Flashes: ", sessionData.Flashes)

  session.Values["sessiondata"] = ggsession.NewSessionData()
  sessionData.AddFlash("message", "Successfully logged out.")
  sessionData.Authenticated = false
  session.Values["sessiondata"] = sessionData
  fmt.Println("Sessin before Redirect (logout): ", session)
  session.Save(r, w)
  http.Redirect(w, r, "/", http.StatusSeeOther)
}
