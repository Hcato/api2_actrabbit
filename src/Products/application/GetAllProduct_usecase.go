package application

import "api/src/Products/domain"

type GetAllProduct struct {
	db domain.RProduct
}

func NewGetAllProduct(db domain.RProduct) *GetAllProduct {
	return &GetAllProduct{db: db}
}

func (cp *GetAllProduct) Execute() ([]domain.Product, error) {
	return cp.db.GetAll()
}
