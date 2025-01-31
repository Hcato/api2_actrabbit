package domain

type RUser interface {
	Save(name string, lastname string) error
	GetAll() ([]User, error)
	Update(id int32, name string, lastname string) error
	Delete(id int32) error
}
