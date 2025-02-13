package application

import "api/src/Users/domain"

type GetByIdUser struct {
	db domain.RUser
}

func NewGetByIdUser(db domain.RUser) *GetByIdUser {
	return &GetByIdUser{db: db}
}

func (getById *GetByIdUser) Execute(id int32) (domain.User, error) {
	return getById.db.GetById(id)
}
