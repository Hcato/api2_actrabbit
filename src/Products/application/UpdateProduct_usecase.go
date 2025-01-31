package application

import "api/src/Products/domain"

type UpdateProduct struct {
	db domain.RProduct
}

func NewUpdateProduct(db domain.RProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}
func (uc *UpdateProduct) Execute(id int32, name string, price float32) error {
	return uc.db.Update(id, name, price)
}
