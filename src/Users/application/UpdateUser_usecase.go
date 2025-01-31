package application

import "api/src/Users/domain"

type UpdateProduct struct {
	db domain.RUser
}

func NewUpdateProduct(db domain.RUser) *UpdateProduct {
	return &UpdateProduct{db: db}
}
func (uc *UpdateProduct) Execute(id int32, name string, lastname string) error {
	return uc.db.Update(id, name, lastname)
}
