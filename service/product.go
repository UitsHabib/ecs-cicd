package service

type Product struct {
	Name     string  `json:"name"`
	ShopName string  `json:"shop_name"`
	Price    float64 `json:"price"`
}
