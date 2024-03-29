package products

import (
	"ducco/microservices/ducco_products/repository/categories"
)

type ProductPipe struct {
	CategoryId     *uint32 `json:"categoryId"`
	NameFO         *string `json:"nameFO"`
	CardTitleFO    *string `json:"cardTitleFO"`
	CardSubTitleFO *string `json:"cardSubTitleFO"`
	CardImgUrlFO   *string `json:"cardImgUrlFO"`
	CardShowFO     *uint8  `json:"cardShowFO"`
	CardOrderFO    *uint8  `json:"cardOrderFO"`
	InsTimestamp   *uint64 `json:"insTimestamp"`
}

func ItemsCustomer(categories []categories.Category) []ProductPipe {
	categoriesPipe := []ProductPipe{}

	for _, category := range categories {
		categoriesPipe = append(categoriesPipe, ProductPipe{
			CategoryId:     category.CategoryId,
			NameFO:         category.NameFO,
			CardTitleFO:    category.CardTitleFO,
			CardSubTitleFO: category.CardSubTitleFO,
			CardImgUrlFO:   category.CardImgUrlFO,
			CardShowFO:     category.CardShowFO,
			CardOrderFO:    category.CardOrderFO,
			InsTimestamp:   category.InsTimestamp,
		})
	}

	return categoriesPipe
}
