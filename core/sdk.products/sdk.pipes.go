package sdk_products

import (
	"ducco/microservices/ducco_products/repository/products"
)

//+ Products
type ProductsSearchItemDataIn struct {
	ProductId uint32
}

type ProductsSearchItemDataOut struct {
	Success bool             `json:"success"`
	Product products.Product `json:"product"`
}
