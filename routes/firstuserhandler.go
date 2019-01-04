package routes

import (
  "fmt"
  "net/http"
  "html/template"

  "github.com/cagox/gge/ggsession"
  "github.com/cagox/gge/config"
  "github.com/cagox/gge/models/user"

)

type firstAdminData struct {
  ggsession.BasePageData
  IsTokenSet bool
  Email      string
  Name       string
}


func firstUserHandle(w http.ResponseWriter, r *http.Request) {
  session := ggsession.GetSession(w, r)
  sessionData := ggsession.GetSessionData(session)
  pageData := firstAdminData{}
  pageData.BasicData(sessionData)
  pageData.Page="FirstUser"
  session.Save(r,w)

  //Make sure the database is actually empty and that they didn't come to this page on accident.
  if (user.AreThereAnyUsers()) {
    sessionData.AddFlash("error", "The Database is not empty.")
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  t := template.New("base.html")
  t, err := t.ParseFiles(config.Config.TemplateRoot+"/base/base.html", config.Config.TemplateRoot+"/admin/firstadmin.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  if config.Config.AdminToken != "" {
    pageData.IsTokenSet = true
  }


  if (r.Method == "GET") {
    pageData.Flashes = sessionData.GetFlashes(true)
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    t.Execute(w, pageData)
    return
  }


  //If method == POST, we start processing the form.
  r.ParseForm()
  newUser := user.Form{}
  newUser.Email = r.FormValue("email")
  newUser.Password = r.FormValue("password")
  newUser.Name = r.FormValue("name")
  adminToken := r.FormValue("admintoken")


  if (adminToken != config.Config.AdminToken) {
    sessionData.AddFlash("error", "Admin Token Does Not Match!")
    pageData.Flashes = sessionData.GetFlashes(true)
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    t.Execute(w, pageData)
    return
  }


  errors := user.ValidateUserForm(newUser, true)
  if len(errors) != 0 {
    for i := range errors {
      sessionData.AddFlash("error", errors[i])
    }
    pageData.Flashes = sessionData.GetFlashes(true)
    session.Values["sessiondata"] = sessionData
    session.Save(r,w)
    t.Execute(w, pageData)
    return
  }

  createdUser := user.CreateUserFromForm(newUser)

  createdUser.SetPassword(r.FormValue("password"))
  createdUser.IsAdmin = true

  mongoSession := config.Config.MongoSession.Clone()
  defer mongoSession.Close()
  collection :=  mongoSession.DB("gge").C("users")

  err = collection.Insert(&createdUser)

  if err != nil {
    //TODO: Proper Erorr handling
    fmt.Println(err)
  }

  //config.Config.Database.Create(&createdUser)
  //profile.UserID = createdUser.ID
  //config.Config.Database.Create(&profile)

  sessionData.AddFlash("message", "User Successfully Created")
  session.Values["sessiondata"] = sessionData
  session.Save(r,w)
  http.Redirect(w, r, "/", http.StatusSeeOther)
  return
}
