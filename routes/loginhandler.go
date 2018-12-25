package routes

import (
  "net/http"
  "fmt"

  //"github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/models/user"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w,r)
  sessionData := ggsession.GetSessionData(session)
  fmt.Println("Login Handler Flashes: ", sessionData.Flashes)

  if sessionData.Authenticated {
    fmt.Println("Login Error: Already Authenticated.")
    sessionData.AddFlash("message", "Already Logged In.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  if (r.Method == "POST") {

    r.ParseForm()
    newUser := user.GetUserByEmail(r.FormValue("email"))
    if (newUser.Authenticate(r.FormValue("password"))) {
      sessionData.AddFlash("message", "We are logging you in!")
      fmt.Println("Flashes after Add(login):", sessionData.Flashes)
      sessionData.Authenticated = true
      sessionData.UserID = newUser.ID
      session.Values["sessiondata"] = sessionData
      fmt.Println("Sessin before Redirect (login): ", session)
      fmt.Println("SessinData before Redirect (login): ", sessionData)
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
