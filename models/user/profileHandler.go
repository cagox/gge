package user

import (
  "net/http"
  "html/template"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
)

type profilePageData struct {
  ggsession.BasePageData

  Name  string

}

func profileHandler(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := profilePageData{}
  pageData.Page = "Profile"
  pageData.Flashes = sessionData.GetFlashes(true)
  pageData.Authenticated = sessionData.Authenticated




  if !sessionData.Authenticated {
    sessionData.AddFlash("error", "You must be logged in to access your profile.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  profile := Profile{}
  if err := config.Config.Database.Where("user_id = ?", sessionData.UserID).First(&profile).Error; err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  currentUser := User{}
  if err := config.Config.Database.Where("id = ?", sessionData.UserID).First(&currentUser).Error; err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }


  pageData.Name = profile.Name
  pageData.IsAdmin = currentUser.IsAdmin


  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/user/profile.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  t.Execute(w, pageData)
  return
}
