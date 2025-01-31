package application

import "api/src/Products/domain"

type CreateProductUsecase struct {
	db domain.RProduct
}

func NewCreateProduct(db domain.RProduct) *CreateProductUsecase {
	return &CreateProductUsecase{db: db}
}

func (cp *CreateProductUsecase) Execute(name string, price float32) error {
	return cp.db.Save(name, price)
}
