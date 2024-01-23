package register

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRegister(r *mux.Router) {
  prefix := "/api"

  r.HandleFunc(prefix + "/register", registerUser).Methods(http.MethodPost)
  r.HandleFunc(prefix + "/login", userLogin).Methods(http.MethodPost)
}
