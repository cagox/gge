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

    if !user.IsEmailValidated(newUser.Email) {
      sessionData.Authenticated = false
      sessionData.AddFlash("error", "Please Validate Your Email.")
      session.Values["sessiondata"] = sessionData
      session.Save(r,w)
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }

    if (newUser.Authenticate(r.FormValue("password"))) {
      sessionData.AddFlash("success", "We are logging you in!")

      sessionData.Authenticated = true
      sessionData.Email = newUser.Email
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


//TODO: Add code to redirec users without verified email addresses to a page that will let them resend the verification email.
