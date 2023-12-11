package products

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model         `gorm:"table:Products"`
	ProductId          *string `gorm:"column:productId"`
	Name               *string `gorm:"column:name"`
	Description        *string `gorm:"column:description"`
	NameFO             *string `gorm:"column:nameFO"`
	DescriptionFO      *string `gorm:"column:descriptionFO"`
	ImageUrlCardFO     *string `gorm:"column:imageUrlCardFO"`
	DetailDocIdFO      *string `gorm:"column:detailDocIdFO"`
	DetailImagesUrlsFO *string `gorm:"column:detailImagesUrlsFO"`
	Stock              *uint32 `gorm:"column:stock"`
	Price              *uint32 `gorm:"column:price"`
	InsTimestamp       *uint64 `gorm:"column:insTimestamp"`
}

type ItemsDBIn struct {
	StockGTE *uint32 `pattern:">" column:"stock"`
}

type ProductsWhere struct {
	ProductIdIn *[]uint32 `pattern:"IN" db:"productId"`
	StockGTE    *uint32   `pattern:"=" db:"stock"`
}
