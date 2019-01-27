package ggsession


import (
  "net/http"
)



func MustBeAdmin(handler http.HandlerFunc) http.HandlerFunc {

  return func(w http.ResponseWriter, r *http.Request) {
    session := GetSession(w,r)
    sessionData := GetSessionData(session)
    pageData := BasePageData{}
    pageData.BasicData(sessionData)

    if !sessionData.Authenticated {
      sessionData.AddFlash("error", "You must be logged in to access that page.")
      session.Values["sessiondata"] = sessionData
      session.Save(r,w)
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }
    if !pageData.IsAdmin {
      sessionData.AddFlash("error", "You must be an administrator to access that page.")
      session.Values["sessiondata"] = sessionData
      session.Save(r,w)
      http.Redirect(w, r, "/", http.StatusSeeOther)
      return
    }
    handler(w,r)
  }





}
