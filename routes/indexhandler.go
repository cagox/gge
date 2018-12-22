package routes

import (
  "fmt"
  "net/http"
  "html/template"

  "github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  //"github.com/cagox/gge/apps/user"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  //TODO: Test for empty database.
  //If database is Empty, go to Admin User Creation Screen.
  //users := user.GetUsers()
  //if (len(users)==0){
  //  http.Redirect(w, r, "/users/firstuser", http.StatusSeeOther)
  //}

  session, err := ggsession.Store.Get(r, "gge-cookie")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  data := session.Values["pagedata"]

  var page = &ggsession.PageData{}
  var pageData = ggsession.PageData{}
  var ok bool

  if data != nil {
    if page, ok = data.(*ggsession.PageData); !ok {
      //We will assume that the session is brand new.
      fmt.Println("No PageData")
      pageData.UserName = ""
      pageData.Email = ""
      pageData.Authenticated = false
    } else {
      fmt.Println("User: "+page.UserName+" "+page.Email)
      pageData = *page
    }
  } else {
    fmt.Println("Data Nil")
    pageData.UserName = ""
    pageData.Email = ""
    pageData.Authenticated = false
  }

  t := template.New("base.html")
  t, err = t.ParseFiles(config.Config.TemplateRoot+"/base/base.html")
  if err != nil {
    fmt.Println(err.Error())
  }


  session.Save(r,w)
  t.Execute(w, pageData)

}
