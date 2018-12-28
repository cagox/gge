package routes

import (
  "net/http"
  //"fmt"

  //"github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/models/user"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w,r)
  sessionData := ggsession.GetSessionData(session)

  if sessionData.Authenticated {
    sessionData.AddFlash("info", "Already Logged In.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  if (r.Method == "POST") {

    r.ParseForm()
    newUser := user.GetUserByEmail(r.FormValue("email"))
    if (newUser.Authenticate(r.FormValue("password"))) {
      sessionData.AddFlash("success", "We are logging you in!")

      sessionData.Authenticated = true
      sessionData.UserID = newUser.ID
      session.Values["sessiondata"] = sessionData

      session.Save(r,w)
    } else {
      sessionData.Authenticated = false
      sessionData.AddFlash("error", "Email or Password did not match.")
      session.Values["sessiondata"] = sessionData
      session.Save(r,w)
    }
  }

  http.Redirect(w, r, "/", http.StatusSeeOther)

}
