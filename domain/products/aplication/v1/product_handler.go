package v1

import (
	"Reactproject-backend_inGoland/domain/products/domain/model"
	repoDomain "Reactproject-backend_inGoland/domain/products/domain/repository"
	"Reactproject-backend_inGoland/domain/products/infrastructure/persistence"
	"Reactproject-backend_inGoland/infrastructure/database"
	"Reactproject-backend_inGoland/infrastructure/middleware"
	"fmt"
	"net/http"
)

//ProductRouter rutas
type ProductRouter struct {
	Repo repoDomain.ProductRepository
}

//NewProductHandler aa
func NewProductHandler(db *database.Data) *ProductRouter {
	return &ProductRouter{
		Repo: persistence.NewProductRepository(db),
	}
}

//CreateProductHandler aaa
func (prod *ProductRouter) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	ctx := r.Context()
	result, err := prod.Repo.CreateProductHandler(ctx, &product)

	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusConflict, "conflict", err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), result))
	_ = middleware.JSON(w, r, http.StatusCreated, result)
}
