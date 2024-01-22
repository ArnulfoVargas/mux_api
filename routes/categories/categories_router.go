package category_routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

func HandleCategories(r *mux.Router) {
	prefix := "/api/gorm"

	r.HandleFunc(prefix+"/category", getCategory).Methods(http.MethodGet)
	r.HandleFunc(prefix+"/category/{id:[0-9]+}", getCategoryByID).Methods(http.MethodGet)
	r.HandleFunc(prefix+"/category", postCategory).Methods(http.MethodPost)
	r.HandleFunc(prefix+"/category/{id:[0-9]+}", putCategory).Methods(http.MethodPut)
	r.HandleFunc(prefix+"/category/{id:[0-9]+}", deleteCategory).Methods(http.MethodDelete)
}
