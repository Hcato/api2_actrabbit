package application

import "api/src/Users/domain"

type DeleteUserUsecase struct {
	db domain.RUser
}

func NewDeleteUser(db domain.RUser) *DeleteUserUsecase {
	return &DeleteUserUsecase{db: db}

}

func (uc *DeleteUserUsecase) Execute(id int32) error {
	return uc.db.Delete(id)
}
