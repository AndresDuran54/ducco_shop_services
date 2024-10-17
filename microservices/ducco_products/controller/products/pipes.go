package products

import (
	"ducco/microservices/ducco_products/repository/products"
	"encoding/json"
)

type ProductPipe struct {
	ProductId              *uint32        `json:"productId,omitempty"`
	Name                   *string        `json:"name,omitempty"`
	Description            *string        `json:"description,omitempty"`
	NameFO                 *string        `json:"nameFO,omitempty"`
	DescriptionFO          *string        `json:"descriptionFO,omitempty"`
	CardTitleFO            *string        `json:"cardTitleFO,omitempty"`
	CardSubTitleFO         *string        `json:"cardSubTitleFO,omitempty"`
	CardImgUrlFO           *string        `json:"cardImgUrlFO,omitempty"`
	DetailTitleFO          *string        `json:"detailTitleFO,omitempty"`
	DetailSubTitleFO       *string        `json:"detailSubTitleFO,omitempty"`
	DetailDescriptionFO    *string        `json:"detailDescriptionFO,omitempty"`
	DetailImagesUrlsFO     *[]string      `json:"detailImagesUrlsFO,omitempty"`
	DetailDocIdFO          *string        `json:"detailDocIdFO,omitempty"`
	DetailFeaturesFO       *[]interface{} `json:"detailFeaturesFO,omitempty"`
	InventoryStock         *uint32        `json:"inventoryStock,omitempty"`
	InventorySalesQuantity *uint32        `json:"inventorySalesQuantity,omitempty"`
	InventoryPrice         *uint32        `json:"inventoryPrice,omitempty"`
	InsTimestamp           *uint64        `json:"insTimestamp,omitempty"`
}

func ItemsCustomer(products []products.Product) []ProductPipe {
	productsPipe := []ProductPipe{}

	for _, product := range products {
		detailImagesUrlsFO := &[]string{}
		detailFeaturesFO := &[]interface{}{}

		json.Unmarshal([]byte(*product.DetailImagesUrlsFO), detailImagesUrlsFO)
		json.Unmarshal([]byte(*product.DetailFeaturesFO), detailFeaturesFO)

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
			DetailFeaturesFO:       detailFeaturesFO,
			InventoryStock:         product.InventoryStock,
			InventorySalesQuantity: product.InventorySalesQuantity,
			InventoryPrice:         product.InventoryPrice,
			InsTimestamp:           product.InsTimestamp,
		})
	}

	return productsPipe
}

func ItemCustomer(product *products.Product) ProductPipe {
	productPipe := ProductPipe{}

	detailImagesUrlsFO := &[]string{}
	detailFeaturesFO := &[]interface{}{}

	json.Unmarshal([]byte(*product.DetailImagesUrlsFO), detailImagesUrlsFO)
	json.Unmarshal([]byte(*product.DetailFeaturesFO), detailFeaturesFO)

	productPipe = ProductPipe{
		ProductId:           product.ProductId,
		DetailTitleFO:       product.DetailTitleFO,
		DetailSubTitleFO:    product.DetailSubTitleFO,
		DetailDescriptionFO: product.DetailDescriptionFO,
		DetailImagesUrlsFO:  detailImagesUrlsFO,
		DetailDocIdFO:       product.DetailDocIdFO,
		DetailFeaturesFO:    detailFeaturesFO,
		InventoryPrice:      product.InventoryPrice,
	}

	return productPipe
}
