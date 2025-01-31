package application

import "api/src/Products/domain"

type DeleteProductUsecase struct {
	db domain.RProduct
}

func NewDeleteProduct(db domain.RProduct) *DeleteProductUsecase {
	return &DeleteProductUsecase{db: db}

}

func (uc *DeleteProductUsecase) Execute(id int32) error {
	return uc.db.Delete(id)
}
