package crypto

import (
  //"fmt"
  "log"

  "golang.org/x/crypto/bcrypt"
)

//HashPassword returns a hashed version of the password.
func HashPassword(password string) string {
  hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    //TODO: Do proper error handling
    log.Fatal(err)
  }
  return string(hash)
}

//ComparePassword tells us if the password matches
func ComparePassword(password string, databaseHash string) bool {
  return databaseHash == string(HashPassword(password))
}
