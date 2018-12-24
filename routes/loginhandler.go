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
    fmt.Println("Parsing Loging Form")
    fmt.Println("Email: "+r.FormValue("email"))
    fmt.Println("Password: "+r.FormValue("password"))

    r.ParseForm()
    newUser := user.GetUserByEmail(r.FormValue("email"))
    fmt.Println("Tried to get User: ", newUser)
    if (newUser.Authenticate(r.FormValue("password"))) {
      sessionData.AddFlash("message", "We are logging you in!")
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
