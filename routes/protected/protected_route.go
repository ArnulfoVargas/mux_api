package protected

import (
	"encoding/json"
	"muxapi/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleProtected(r *mux.Router) {
  r.HandleFunc("/protected", middleware.ValidateJWT(protectedData))
}

func protectedData(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(map[string]any{
    "status" : http.StatusOK,
    "message": "This message is protected",
  })
}
