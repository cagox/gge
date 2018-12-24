package ggsession
import (
  "encoding/gob"
  //"github.com/gorilla/securecookie"
  "github.com/gorilla/sessions"
  //"fmt"
  "github.com/cagox/gge/config"
)

//Store is the session store.
var Store *sessions.CookieStore

func init() {
  //authKeyOne := securecookie.GenerateRandomKey(64)
  //encryptionKeyOne := securecookie.GenerateRandomKey(32)

  authKeyOne := []byte(config.Config.AuthKey)
  encryptionKeyOne := []byte(config.Config.EncKey)

  Store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)

  Store.Options = &sessions.Options{  //Max Age 30 Days. This site is not exactly high risk.
    Path:   "/",
    MaxAge: 3600 * 24 * 30,
    HttpOnly: true,
  }

  //Register necessary structs.
  gob.Register(PageData{})

}

//PageData is a struct of information to pass to sessions and pages.
type PageData struct {
  UserName      string
  Email         string
  Authenticated bool
}

//ZeroPageData returns a zeroed out PageData struct. Most importantly Authenticated is set to false.
func ZeroPageData() PageData {
  return PageData{UserName: "", Email: "", Authenticated: false}
}

//GetPageData grabs the PageData struct from the cookie and returns it.
func GetPageData(session *sessions.Session) PageData {
  data := session.Values["pagedata"]

  var pageData = PageData{}

  if data != nil {
    if page, ok := data.(PageData); !ok {
      //We will assume that the session is brand new.
      pageData.UserName = ""
      pageData.Email = ""
      pageData.Authenticated = false
    } else {
      pageData = page
    }
  } else {
    pageData.UserName = ""
    pageData.Email = ""
    pageData.Authenticated = false
  }


  return pageData
}
