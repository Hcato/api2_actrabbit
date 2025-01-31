package domain

type Product struct {
	Id    int32
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{Id: 1, Name: name, Price: price}
}

func (p *Product) SaveName(name string) {
	p.Name = name
}
