package routes

import (
  "net/http"

  //"github.com/cagox/gge/config"
  "github.com/cagox/gge/ggsession"
  //"github.com/cagox/gge/apps/user"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
  session, err := ggsession.Store.Get(r, "gge-cookie")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  pageData := ggsession.GetPageData(session)

  if pageData.Authenticated {
    session.AddFlash("Already Logged In.")
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }

  if (r.Method == "POST") {
    r.ParseForm()
    pageData.Email = r.FormValue("email")
    pageData.UserName = r.FormValue("email")
    pageData.Authenticated = true
    session.Values["pagedata"] = pageData
    session.Save(r,w)
  }
  http.Redirect(w, r, "/", http.StatusSeeOther)

}
