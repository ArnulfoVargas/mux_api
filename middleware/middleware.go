package middleware

import (
	"fmt"
	"muxapi/database"
	"muxapi/models"
	"muxapi/utils"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    err := godotenv.Load()
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      utils.EncodeGenericResponse(w, "Error", "Unexpected Error")
      return
    }

    jwtSignature := []byte(os.Getenv("JWT_SIGNATURE"))
    authHeader := r.Header.Get("Authorization")

    if len(authHeader) == 0 {
      w.WriteHeader(http.StatusUnauthorized)
      utils.EncodeGenericResponse(w, "Error", "Unauthorized")
      return
    }

    splitBearer := strings.Split(authHeader, " ")
    if len(splitBearer) != 2 {
      w.WriteHeader(http.StatusUnauthorized)
      utils.EncodeGenericResponse(w, "Error", "Unauthorized")
      return
    }

    splitToken := strings.Split(splitBearer[1], ".")
    if len(splitToken) != 3 {
      w.WriteHeader(http.StatusUnauthorized)
      utils.EncodeGenericResponse(w, "Error", "Unauthorized")
      return
    }

    tk := strings.TrimSpace(splitBearer[1])

    token, err := jwt.Parse(tk, func (tok *jwt.Token) (interface{}, error) {
      if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected Login Error")
      }

      return jwtSignature, nil
    })

    if err != nil {
      w.WriteHeader(http.StatusUnauthorized)
      utils.EncodeGenericResponse(w, "Error", "Unauthorized")
      return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
      w.WriteHeader(http.StatusUnauthorized)
      utils.EncodeGenericResponse(w, "Error", "Unauthorized")
      return
    } else {
      user := models.User{}
      if err := database.Database.Where("mail=?", claims["mail"]).First(&user); err.Error != nil {
        w.WriteHeader(http.StatusUnauthorized)
        utils.EncodeGenericResponse(w, "Error", "Unauthorized")
        return
      } 

      next.ServeHTTP(w, r)
    }
  }
}
