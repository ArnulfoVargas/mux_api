package main

import (
	"fmt"
	"muxapi/routes"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()

  if err != nil {
    panic("Couldnt load .env file")
  }

  address := generateAddress()
  fmt.Println("Running server at " + address)
  panic(http.ListenAndServe(address, routes.HandleRoutes()))
}

func generateAddress() string {
  sb := strings.Builder{}

  sb.WriteString(os.Getenv("ADDRESS"))
  sb.WriteByte(':')
  sb.WriteString(os.Getenv("PORT"))
  
  return sb.String()
}
