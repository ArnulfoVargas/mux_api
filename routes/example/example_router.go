package example_routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleExample(mux *mux.Router) {
	prefix := "/api/example"

	mux.HandleFunc(prefix+"/", exampleGet).Methods(http.MethodGet)
	mux.HandleFunc(prefix+"/", examplePost).Methods(http.MethodPost)
	mux.HandleFunc(prefix+"/{id:[0-9]+}", examplePut).Methods(http.MethodPut)
	mux.HandleFunc(prefix+"/{id:[0-9]+}", exampleDelete).Methods(http.MethodDelete)
	mux.HandleFunc(prefix+"/{id:[0-9]+}", exampleParamsGet).Methods(http.MethodGet)
	mux.HandleFunc(prefix+"/query", exampleQueryGet).Methods(http.MethodGet)
	mux.HandleFunc(prefix+"/file", exampleFileUpload).Methods(http.MethodPost)
	mux.HandleFunc(prefix+"/file", exampleFilePreview).Methods(http.MethodGet)
}
