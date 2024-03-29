package products

import (
	"ducco/microservices/ducco_products/repository/products"
	"encoding/json"
)

type ProductPipe struct {
	ProductId              *uint32   `json:"productId"`
	Name                   *string   `json:"name"`
	Description            *string   `json:"description"`
	NameFO                 *string   `json:"nameFO"`
	DescriptionFO          *string   `json:"descriptionFO"`
	CardTitleFO            *string   `json:"cardTitleFO"`
	CardSubTitleFO         *string   `json:"cardSubTitleFO"`
	CardImgUrlFO           *string   `json:"cardImgUrlFO"`
	DetailTitleFO          *string   `json:"detailTitleFO"`
	DetailSubTitleFO       *string   `json:"detailSubTitleFO"`
	DetailDescriptionFO    *string   `json:"detailDescriptionFO"`
	DetailImagesUrlsFO     *[]string `json:"detailImagesUrlsFO"`
	DetailDocIdFO          *string   `json:"detailDocIdFO"`
	InventoryStock         *uint32   `json:"inventoryStock"`
	InventorySalesQuantity *uint32   `json:"inventorySalesQuantity"`
	InventoryPrice         *uint32   `json:"inventoryPrice"`
	InsTimestamp           *uint64   `json:"insTimestamp"`
}

func ItemsCustomer(products []products.Product) []ProductPipe {
	productsPipe := []ProductPipe{}

	for _, product := range products {
		detailImagesUrlsFO := &[]string{}

		json.Unmarshal([]byte(*product.DetailImagesUrlsFO), detailImagesUrlsFO)

		productsPipe = append(productsPipe, ProductPipe{
			ProductId:              product.ProductId,
			Name:                   product.Name,
			Description:            product.Description,
			NameFO:                 product.NameFO,
			DescriptionFO:          product.DescriptionFO,
			CardTitleFO:            product.CardTitleFO,
			CardSubTitleFO:         product.CardSubTitleFO,
			CardImgUrlFO:           product.CardImgUrlFO,
			DetailTitleFO:          product.DetailTitleFO,
			DetailSubTitleFO:       product.DetailSubTitleFO,
			DetailDescriptionFO:    product.DetailDescriptionFO,
			DetailImagesUrlsFO:     detailImagesUrlsFO,
			DetailDocIdFO:          product.DetailDocIdFO,
			InventoryStock:         product.InventoryStock,
			InventorySalesQuantity: product.InventorySalesQuantity,
			InventoryPrice:         product.InventoryPrice,
			InsTimestamp:           product.InsTimestamp,
		})
	}

	return productsPipe
}
