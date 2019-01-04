package user

import (
  "github.com/cagox/gge/config"
  "regexp"
  "strconv"
)


/*
ValidateUserForm validates the data provided and returns a slice of strings if there are any errors.
Its first argument is a filled in UserForm struct. This is the data to be validated/
Its second argument is a bool value declaring if the form is for a new user rather than just updating a user.
*/
func ValidateUserForm(newUser Form, isNew bool) []string {
  errors := make([]string, 0)

  //Validate Password:
  passwordErrors := ValidatePassword(newUser.Password)
  if passwordErrors != nil {
    for err := range passwordErrors {
      errors = append(errors, passwordErrors[err])
    }
  }

  //Validate the email address format.
  emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  if (!emailRegexp.MatchString(newUser.Email)) {
    errors = append(errors, newUser.Email+" does not appear to be a valid email address.")
  }

  //Validate Email Address is unique:
  if isNew {
    //config.Config.Database.Where("email = ?", newUser.Email).Find(&users)
    if (!isEmailUnique(newUser.Email)) {
      errors = append(errors, "The email address "+newUser.Email+" already exists.")
    }
  }

  //Validate Display name
  if len(newUser.Name) < config.Config.MinimumNameLength {
    errors = append(errors, "The Display Name must be at least "+strconv.Itoa(config.Config.MinimumNameLength)+" characters long.")
  }

  return errors
}

//ValidatePassword validates the offered user password.
func ValidatePassword(password string) []string {
  errors := make([]string, 0)
  if (len(password) < config.Config.MinPasswordLength) {
    errors = append(errors, "The password must be at least "+ strconv.Itoa(config.Config.MinPasswordLength)+" characters.")
  }
  return errors
}

func isEmailUnique(email string) bool {

  return true
}
