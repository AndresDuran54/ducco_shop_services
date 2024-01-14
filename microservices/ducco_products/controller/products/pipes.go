package products

import (
	"ducco/microservices/ducco_products/repository/products"
	"encoding/json"
)

type ProductPipe struct {
	ProductId          *uint32   `json:"productId"`
	Name               *string   `json:"name"`
	Description        *string   `json:"description"`
	NameFO             *string   `json:"nameFO"`
	DescriptionFO      *string   `json:"descriptionFO"`
	ImageUrlCardFO     *string   `json:"imageUrlCardFO"`
	DetailDocIdFO      *string   `json:"detailDocIdFO"`
	DetailImagesUrlsFO *[]string `json:"detailImagesUrlsFO"`
	Stock              *uint32   `json:"stock"`
	Price              *uint32   `json:"price"`
	InsTimestamp       *uint64   `json:"insTimestamp"`
}

func ItemsCustomer(products []products.Product) []ProductPipe {
	productsPipe := []ProductPipe{}

	for _, product := range products {
		detailImagesUrlsFO := &[]string{}

		json.Unmarshal([]byte(*product.DetailImagesUrlsFO), detailImagesUrlsFO)

		productsPipe = append(productsPipe, ProductPipe{
			ProductId:          product.ProductId,
			Name:               product.Name,
			Description:        product.Description,
			NameFO:             product.NameFO,
			DescriptionFO:      product.DescriptionFO,
			ImageUrlCardFO:     product.ImageUrlCardFO,
			DetailDocIdFO:      product.DetailDocIdFO,
			DetailImagesUrlsFO: detailImagesUrlsFO,
			Stock:              product.Stock,
			Price:              product.Price,
			InsTimestamp:       product.InsTimestamp,
		})
	}

	return productsPipe
}
