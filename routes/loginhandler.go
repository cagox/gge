package routes

import (
  "fmt"
  "net/http"
  //"html/template"

  //"github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  //"github.com/cagox/gge/apps/user"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Login Handler Running")
  session, err := ggsession.Store.Get(r, "gge-cookie")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Println(session)

  var isLoggedIn = false

  data := session.Values["pagedata"]
  var page = &ggsession.PageData{}
  var pageData = &ggsession.PageData{}
  var ok bool

  if page, ok = data.(*ggsession.PageData); ok {
    isLoggedIn = page.Authenticated
  }

  if isLoggedIn {
    fmt.Println("Logged In!")
    session.AddFlash("Already Logged In.")
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  if (r.Method == "POST") {
    r.ParseForm()
    pageData.Email = r.FormValue("email")
    pageData.UserName = r.FormValue("email")
    fmt.Println("Email: "+r.FormValue("email"))
    pageData.Authenticated = true
    session.Values["pagedata"] = pageData
    session.Save(r,w)
    fmt.Println("pageData.Email = "+pageData.Email)
    fmt.Println(session)
  }
  http.Redirect(w, r, "/", http.StatusSeeOther)

}
