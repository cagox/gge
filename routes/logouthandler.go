package routes

import (
  "net/http"
  //"github.com/gorilla/sessions"

  "github.com/cagox/gge/ggsession"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
  session, err := ggsession.Store.Get(r, "gge-cookie")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  pageData := ggsession.ZeroPageData()
  session.Values["pagedata"] = pageData

  session.Save(r, w)
  http.Redirect(w, r, "/", http.StatusSeeOther)
}
