package userhandlers

import(
  "net/http"
  "html/template"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/email"
  "github.com/cagox/gge/models/user"
)



type invitePageData struct {
  ggsession.BasePageData
  Email string
  Token string
}

type emailTemplateData struct {
  Sender string
  Token  string
}



func inviteUserHandler(w http.ResponseWriter, r *http.Request){
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := invitePageData{}
  pageData.Page = "InviteUser"
  pageData.BasicData(sessionData)


  //Setting up the template. The only way this won't get used is if the entire
  //thing crashes anyway.
  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/user/inviteuser.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if r.Method == "POST" {
    bad := false
    r.ParseForm()
    inviteEmail := r.FormValue("email")

    if !user.IsEmailUnique(inviteEmail) {
      sessionData.AddFlash("error", "The User "+inviteEmail+" is already in the database.")
      bad = true
    }

    /*
      I am checking if it is bad first because I am pretty sure checking a boolean
      is less resource intensive than accessing the database.
    */
    if ! bad {
      if user.InviteEmailExists(inviteEmail){
        sessionData.AddFlash("error", "An Invite For "+inviteEmail+" is already in the database.")
        bad = true
      }
    }

    // if bad == true, then the email address is already in the database as
    // either a user or an invite.
    if bad {
      pageData.Flashes = sessionData.GetFlashes(true)
      session.Values["sessiondata"] = sessionData
      session.Save(r,w)
      t.Execute(w, pageData)
      return
    }

    currentUser := user.GetUserByEmail(sessionData.Email)

    invite := user.NewInvitation(inviteEmail)
    invite.Token = user.GenerateInviteToken(64)

    emailData         := email.Data{}  //Data to send to the mailer
    templateData := emailTemplateData{Sender: currentUser.Profile.Name} //Data to send to the template processor.
    templateData.Token = invite.Token

    emailData.To = invite.Email
    emailData.Subject = "Invitation to The God Game Engine Play By Post Site"
    emailData.Body, err = email.ParseHTMLEmail(config.Config.TemplateRoot+"/email/invite.html", templateData)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    emailData.HTML = true



    err = email.SystemEmail(emailData)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    sessionData.AddFlash("success", "Invitation Email Sent!")

    err = user.InsertInvitation(invite)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    sessionData.AddFlash("success", "Invitation Successfully Inserted")

    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }

  pageData.Flashes = sessionData.GetFlashes(true)
  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  t.Execute(w, pageData)
  return
}
