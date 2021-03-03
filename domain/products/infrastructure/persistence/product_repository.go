package persistence

import (
	"Reactproject-backend_inGoland/domain/products/aplication/v1/response"
	"Reactproject-backend_inGoland/domain/products/domain/model"
	repoDomain "Reactproject-backend_inGoland/domain/products/domain/repository"
	"Reactproject-backend_inGoland/infrastructure/database"
	"context"
	"database/sql"
)

type sqlProductRepo struct {
	Conn *database.Data
}

//NewProductRepository aaa
func NewProductRepository(Conn *database.Data) repoDomain.ProductRepository {
	return &sqlProductRepo{
		Conn: Conn,
	}
}

func (sr *sqlProductRepo) CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, insertProduct)
	if err != nil {
		return &response.ProductCreateResponse{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, &product.ProductoID, &product.ProductoNombre, &product.ProductoCantidad, &product.ProductoUsercreacion, &product.ProductoFechCreacion, &product.ProductoUserModificacion, &product.ProductoFechaModificacion)
	var idResult string
	err = row.Scan(&idResult)
	if err != sql.ErrNoRows {
		return &response.ProductCreateResponse{}, err
	}
	ProductResponse := response.ProductCreateResponse{
		Message: "product created",
	}
	return &ProductResponse, nil
}
