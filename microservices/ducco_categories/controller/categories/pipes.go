package products

import (
	"ducco/microservices/ducco_categories/repository/categories"
	"encoding/json"
)

type CardFiltersFO struct {
	Title   string                `json:"title"`
	Filter  string                `json:"filter"`
	Type    string                `json:"type"`
	Options []CardFiltersFOOption `json:"options,omitempty"`
	Pattern string                `json:"pattern"`
}

type CardFiltersFOOption struct {
	ID    string  `json:"id"`
	Type  *string `json:"type,omitempty"`
	Label string  `json:"label"`
}

type CardOrdersFO struct {
	Val   string `json:"val"`
	Label string `json:"label"`
	Order string `json:"order"`
}

type ProductPipe struct {
	CategoryId     *uint32          `json:"categoryId"`
	NameFO         *string          `json:"nameFO"`
	CardTitleFO    *string          `json:"cardTitleFO"`
	CardSubTitleFO *string          `json:"cardSubTitleFO"`
	CardImgUrlFO   *string          `json:"cardImgUrlFO"`
	CardShowFO     *uint8           `json:"cardShowFO"`
	CardOrderFO    *uint8           `json:"cardOrderFO"`
	CardFiltersFO  *[]CardFiltersFO `json:"cardFiltersFO"`
	CardOrdersFO   *[]CardOrdersFO  `json:"cardOrdersFO"`
	InsTimestamp   *uint64          `json:"insTimestamp"`
}

func ItemsCustomer(categories []categories.Category) []ProductPipe {
	categoriesPipe := []ProductPipe{}

	for _, category := range categories {
		cardFiltersFO := &[]CardFiltersFO{}
		cardOrdersFO := &[]CardOrdersFO{}

		json.Unmarshal([]byte(*category.CardFiltersFO), cardFiltersFO)
		json.Unmarshal([]byte(*category.CardOrdersFO), cardOrdersFO)

		categoriesPipe = append(categoriesPipe, ProductPipe{
			CategoryId:     category.CategoryId,
			NameFO:         category.NameFO,
			CardTitleFO:    category.CardTitleFO,
			CardSubTitleFO: category.CardSubTitleFO,
			CardImgUrlFO:   category.CardImgUrlFO,
			CardShowFO:     category.CardShowFO,
			CardOrderFO:    category.CardOrderFO,
			CardFiltersFO:  cardFiltersFO,
			CardOrdersFO:   cardOrdersFO,
			InsTimestamp:   category.InsTimestamp,
		})
	}

	return categoriesPipe
}
