package validations

import (
	"regexp"
	"unicode"
)

var Mail_Regex = regexp.MustCompile("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")

func ValidatePassword(password string) bool {
  var (
    hasMinLen = false
    hasUpper = false
    hasLower = false
    hasNumber = false
  )

  if len(password) >=8 && len(password) <=20 {
    hasMinLen = true
  }

  for _, ch := range password {
    switch {
    case unicode.IsUpper(ch):
      hasUpper = true
    case unicode.IsLower(ch):
      hasLower = true
    case unicode.IsNumber(ch):
      hasNumber = true
    }
  }
  
  return hasMinLen && hasUpper && hasLower && hasNumber
}
