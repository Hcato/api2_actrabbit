package domain

type RProduct interface {
	Save(name string, price float32) error
	GetAll() ([]Product, error)
	GetById(id int32) (Product, error)
	Update(id int32, name string, price float32) error
	Delete(id int32) error
}
