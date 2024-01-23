package register

import (
	"encoding/json"
	"muxapi/database"
	jwt_generator "muxapi/jwt"
	"muxapi/models"
	"muxapi/utils"
	"muxapi/validations"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	var register models.UserDTO
  decoder := json.NewDecoder(r.Body)

  if err := decoder.Decode(&register); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Bad Request")
    return
  }

  if !validations.ValidatePassword(register.Password) {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Password")
    return
  }

  if validations.Mail_Regex.FindStringSubmatch(register.Mail) == nil {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Email")
    return
  }
  
  if len(register.Name) < 2 {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Name")
    return
  }

  if len(register.Phone) < 8 {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Phone")
    return
  }

  user := models.User{}

  if database.Database.Where("mail=?", register.Mail).Find(&user).Limit(1).RowsAffected > 0 {
    w.WriteHeader(http.StatusConflict)
    utils.EncodeGenericResponse(w, "Error", "Email already in use")
    return
  }

  bcost := 8
  bytesPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcost)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Password")
    return
  }

  data := models.User{
    Name: register.Name,
    Password: string(bytesPassword),
    Mail: register.Mail,
    ProfileID: register.ProfileID,
    Phone: register.Phone,
    Date: time.Now(),
  }

  database.Database.Save(&data)
  w.WriteHeader(http.StatusCreated)
  utils.EncodeGenericResponse(w, "Ok", "Registered Successfully")
}

func userLogin(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  loginData := models.LoginDTO{}
  decoder := json.NewDecoder(r.Body)

  if err := decoder.Decode(&loginData); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Password or Email")
    return
  }

  if !validations.ValidatePassword(loginData.Password) {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Password")
    return
  }

  if validations.Mail_Regex.FindStringSubmatch(loginData.Mail) == nil {
    w.WriteHeader(http.StatusBadRequest)
    utils.EncodeGenericResponse(w, "Error", "Invalid Email")
    return
  }

  user := models.User{}

  if database.Database.Where("mail=?", loginData.Mail).Limit(1).Find(&user).RowsAffected > 0 {
    passwordBytes := []byte(loginData.Password)
    passwordEncrypted := []byte(user.Password)
    err := bcrypt.CompareHashAndPassword(passwordEncrypted, passwordBytes)

    if err != nil {
      w.WriteHeader(http.StatusConflict)
      utils.EncodeGenericResponse(w, "Error", "Email or Password are incorrect")
      return
    }

    key, err := jwt_generator.GenerateJWT(user)
    if err != nil {
      w.WriteHeader(http.StatusConflict)
      utils.EncodeGenericResponse(w, "Error", "Error generating the token")
      return
    }

    output := models.LoginResponseDTO{
      Name: user.Name,
      Token: key,
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(output)
  } else {
    w.WriteHeader(http.StatusConflict)
    utils.EncodeGenericResponse(w, "Error", "Email or Password are incorrect")
    return
  }
}
