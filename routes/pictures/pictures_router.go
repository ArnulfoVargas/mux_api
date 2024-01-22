package pictures_routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandlePictures(r *mux.Router) {
  prefix := "/api/gorm"

  r.HandleFunc(prefix + "/pictures/{id:[0-9]+}", postPicture).Methods(http.MethodPost)
  r.HandleFunc(prefix + "/pictures/{id:[0-9]+}", getPictureByProd).Methods(http.MethodGet)
  r.HandleFunc(prefix + "/pictures/{id:[0-9]+}", deletePicture).Methods(http.MethodDelete)
}
