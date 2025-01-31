package application

import "api/src/Users/domain"

type CreateUserUsecase struct {
	db domain.RUser
}

func NewCreateUser(db domain.RUser) *CreateUserUsecase {
	return &CreateUserUsecase{db: db}
}

func (cp *CreateUserUsecase) Execute(name string, lastname string) error {
	return cp.db.Save(name, lastname)
}
