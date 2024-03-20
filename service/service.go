package service

type service struct {
	// all the repos
}

// NewService creates a service with the necessary dependencies
func NewService() Service {
	return &service{}
}

func (s *service) GetUsers() ([]*User, error) {
	usrs := []*User{
		{
			FirstName: "Habibur",
			LastName:  "Rahman",
		},
	}

	return usrs, nil
}

func (s *service) GetShops() ([]*Shop, error) {
	shops := []*Shop{
		{
			Name:      "Pathao",
			OwnerName: "Habibur Rahman",
		},
	}

	return shops, nil
}

func (s *service) GetProducts() ([]*Product, error) {
	products := []*Product{
		{
			Name:     "Salt",
			ShopName: "Pathao",
			Price:    10,
		},
	}

	return products, nil
}

func (s *service) GetBrands() ([]*Brand, error) {
	brands := []*Brand{
		{
			Name:     "Salt",
			ImageURL: "https://pathao.com",
		},
	}

	return brands, nil
}
