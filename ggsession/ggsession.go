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

  authKeyOne := config.Config.AuthKey
  encryptionKeyOne := config.Config.EncKey

  Store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)
  //Store = sessions.NewCookieStore(authKeyOne)

  Store.Options = &sessions.Options{  //Max Age 30 Days. This site is not exactly high risk.
    MaxAge: 3600 * 24 * 30,
    HttpOnly: true,
  }


  //Register necessary structs.
  gob.Register(PageData{})
  //gob.Register(Flash{})

}

//PageData is a struct of information to pass to sessions and pages.
type PageData struct {
  UserName      string
  Email         string
  Authenticated bool
}
