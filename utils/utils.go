package utils

import (
	"encoding/json"
	"net/http"
)

func EncodeGenericResponse(w http.ResponseWriter, status, msg string) {
  json.NewEncoder(w).Encode(map[string]string{
    "status" : status,
    "message": msg,
  })
}
