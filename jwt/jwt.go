package jwt_generator

import (
	"fmt"
	"muxapi/models"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateJWT(user models.User) (string, error){
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
  
  key := []byte(os.Getenv("JWT_SIGNATURE"))

  generatedtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "mail"  : user.Mail,
    "name"  : user.Name,
    "id"    : user.Id,
    "iat"   : time.Now().Unix(),
    "exp"   : time.Now().Add(time.Hour * 24).Unix(),
  })

  jsonToken, err := generatedtoken.SignedString(key) 
  fmt.Println(err)
  return jsonToken, err
}
