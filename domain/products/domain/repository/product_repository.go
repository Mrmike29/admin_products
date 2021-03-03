package repository

import (
	"Reactproject-backend_inGoland/domain/products/aplication/v1/response"
	"Reactproject-backend_inGoland/domain/products/domain/model"
	"context"
)

//ProductRepository type
type ProductRepository interface {
	CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error)
}
