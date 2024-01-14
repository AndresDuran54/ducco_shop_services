package products

type Product struct {
	ProductId          *uint32 `gorm:"column:productId" json:"productId"`
	Name               *string `gorm:"column:name" json:"name"`
	Description        *string `gorm:"column:description" json:"description"`
	NameFO             *string `gorm:"column:nameFO" json:"nameFO"`
	DescriptionFO      *string `gorm:"column:descriptionFO" json:"descriptionFO"`
	ImageUrlCardFO     *string `gorm:"column:imageUrlCardFO" json:"imageUrlCardFO"`
	DetailDocIdFO      *string `gorm:"column:detailDocIdFO" json:"detailDocIdFO"`
	DetailImagesUrlsFO *string `gorm:"column:detailImagesUrlsFO" json:"detailImagesUrlsFO"`
	Stock              *uint32 `gorm:"column:stock" json:"stock"`
	Price              *uint32 `gorm:"column:price" json:"price"`
	InsTimestamp       *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
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
