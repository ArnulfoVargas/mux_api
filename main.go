package main

import (
	"fmt"
	"muxapi/models"
	"muxapi/routes"
	"net/http"
	"os"
	"strings"
)

func main() {
  models.Migrations()

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
