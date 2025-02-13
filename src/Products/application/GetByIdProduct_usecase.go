package application

import "api/src/Products/domain"

type GetByIdProduct struct {
	db domain.RProduct
}

func NewGetByIdProduct(db domain.RProduct) *GetByIdProduct {
	return &GetByIdProduct{db: db}
}

func (getById *GetByIdProduct) Execute(id int32) (domain.Product, error) {
	return getById.db.GetById(id)
}
