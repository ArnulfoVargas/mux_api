package routes

import (
	example_routes "muxapi/routes/example"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRoutes() http.Handler {
	r := mux.NewRouter()

  example_routes.HandleExample(r)

  // CORS
  handler := cors.AllowAll().Handler(r)

  return handler
}
