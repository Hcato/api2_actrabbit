package application

import "api/src/Users/domain"

type GetAllUser struct {
	db domain.RUser
}

func NewGetAllProduct(db domain.RUser) *GetAllUser {
	return &GetAllUser{db: db}
}

func (cp *GetAllUser) Execute() ([]domain.User, error) {
	return cp.db.GetAll()
}
