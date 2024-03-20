package service

type Service interface {
	GetUsers() ([]*User, error)
	GetShops() ([]*Shop, error)
	GetProducts() ([]*Product, error)
	GetBrands() ([]*Brand, error)
}
