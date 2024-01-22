package product_routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleProducts(r *mux.Router) {
  prefix := "/api/gorm"

  r.HandleFunc(prefix + "/products", getProducts).Methods(http.MethodGet)
  r.HandleFunc(prefix + "/products/{id:[0-9]+}", getProductByID).Methods(http.MethodGet)
  r.HandleFunc(prefix + "/products", postProduct).Methods(http.MethodPost)
  r.HandleFunc(prefix + "/products/{id:[0-9]+}", putProduct).Methods(http.MethodPut)
  r.HandleFunc(prefix + "/products", deleteProduct).Methods(http.MethodDelete)
}
