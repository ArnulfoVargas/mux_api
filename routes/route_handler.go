package routes

import (
	"muxapi/routes/categories"
	"muxapi/routes/example"
	"muxapi/routes/pictures"
	"muxapi/routes/products"
	"muxapi/routes/protected"
	"muxapi/routes/register"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRoutes() http.Handler {
	r := mux.NewRouter()

  example_routes.HandleExample(r)
  category_routes.HandleCategories(r)
  product_routes.HandleProducts(r)
  pictures_routes.HandlePictures(r)
  register.HandleRegister(r)
  protected.HandleProtected(r)

  // CORS

  // c := cors.New(cors.Options{
  //   AllowedOrigins: []string{}, <- Allowed sites
  // })

  handler := cors.AllowAll().Handler(r)

  return handler
}
