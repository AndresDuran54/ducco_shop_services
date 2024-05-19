package products

type Product struct {
	ProductId              *uint32 `gorm:"column:productId" json:"productId"`
	Name                   *string `gorm:"column:name" json:"name"`
	Description            *string `gorm:"column:description" json:"description"`
	NameFO                 *string `gorm:"column:nameFO" json:"nameFO"`
	DescriptionFO          *string `gorm:"column:descriptionFO" json:"descriptionFO"`
	CardTitleFO            *string `gorm:"column:cardTitleFO" json:"cardTitleFO"`
	CardSubTitleFO         *string `gorm:"column:cardSubTitleFO" json:"cardSubTitleFO"`
	CardImgUrlFO           *string `gorm:"column:cardImgUrlFO" json:"cardImgUrlFO"`
	DetailTitleFO          *string `gorm:"column:detailTitleFO" json:"detailTitleFO"`
	DetailSubTitleFO       *string `gorm:"column:detailSubTitleFO" json:"detailSubTitleFO"`
	DetailImagesUrlsFO     *string `gorm:"column:detailImagesUrlsFO" json:"detailImagesUrlsFO"`
	DetailDescriptionFO    *string `gorm:"column:detailDescriptionFO" json:"detailDescriptionFO"`
	DetailDocIdFO          *string `gorm:"column:detailDocIdFO" json:"detailDocIdFO"`
	InventoryStock         *uint32 `gorm:"column:inventoryStock" json:"inventoryStock"`
	InventorySalesQuantity *uint32 `gorm:"column:inventorySalesQuantity" json:"inventorySalesQuantity"`
	InventoryPrice         *uint32 `gorm:"column:inventoryPrice" json:"inventoryPrice"`
	InsTimestamp           *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type BuildWhere struct {
	ProductId *uint32 `db:"productId" pattern:"="`
}

type ItemDBIn struct {
	ProductId *uint32 `column:"productId" pattern:">"`
	Label     string
	Trace     string
}

type ItemsDBIn struct {
	StockGTE     *uint32 `pattern:">" column:"stock"`
	OrderCol     *string
	Order        *string
	FilterVals   *string
	OrderVals    *string
	EnablePaging *bool
	PagingSize   *int
	PagingIndex  *int
	Label        string
	Trace        string
}

type ProductsWhere struct {
	ProductIdIn *[]uint32 `pattern:"IN" db:"productId"`
	StockGTE    *uint32   `pattern:"=" db:"stock"`
}
