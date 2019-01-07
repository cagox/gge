package userhandlers

import (
  "net/http"
  "html/template"

  "github.com/gorilla/mux"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/models/user"

)

type acceptInvitePageData struct {
  ggsession.BasePageData
  Email string
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

  if !user.InviteTokenExists(pageData.Token){
    sessionData.AddFlash("error", "No Invitation Matching that Token.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  invite, err := user.InviteByToken(pageData.Token)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  pageData.Email = invite.Email


  t := template.New("base.html")
  t, err = t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/user/acceptinvite.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if r.Method == "POST" {
    userForm := user.Form{}
    r.ParseForm()
    userForm.Email = r.FormValue("email")
    userForm.Name = r.FormValue("name")
    userForm.Password = r.FormValue("password")

    newUser := user.CreateUserFromForm(userForm)
    newUser.EmailIsVerified = true

    mongoSession := config.Config.MongoSession.Clone()
    defer mongoSession.Close()

    users := mongoSession.DB("gge").C("users")

    err = users.Insert(&newUser)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    err = user.RemoveInvitation(pageData.Token)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    sessionData.AddFlash("success", "Account created for "+newUser.Email)
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
}


  pageData.Flashes = sessionData.GetFlashes(true)
  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  t.Execute(w, pageData)
  return
}


/*
TODO: Add code to verify that the email address provided matches the invitation.
      If it does not, then we should send an email verification message.
*/
