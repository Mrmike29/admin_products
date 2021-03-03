package infrastructure

import (
	"net/http"

	"Reactproject-backend_inGoland/infrastructure/database"

	"github.com/go-chi/chi"

	v1Product "Reactproject-backend_inGoland/domain/products/aplication/v1"
)

// RoutesProducts aa
func RoutesProducts(conn *database.Data) http.Handler {
	router := chi.NewRouter()

	products := v1Product.NewProductHandler(conn)
	router.Mount("/products", routesProduct(products))

	return router
}

func routesProduct(handler *v1Product.ProductRouter) http.Handler {
	router := chi.NewRouter()
	router.Post("/", handler.CreateProductHandler)
	return router
}
