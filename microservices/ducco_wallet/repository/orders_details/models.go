package orders_details

type OrderDetail struct {
	OrderDetailId *uint32 `gorm:"column:orderDetailId;primaryKey;autoIncrement" json:"orderDetailId"`
	OrderId       *uint32 `gorm:"column:orderId" json:"orderId"`
	ProductId     *uint32 `gorm:"column:productId" json:"productId"`
	Quantity      *uint32 `gorm:"column:quantity" json:"quantity"`
	InsTimestamp  *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type BuildWhere struct {
	OrderDetailId *uint32 `db:"orderDetailId" pattern:"="`
	Token         *string `db:"token" pattern:"="`
}

//+ ITEM
type ItemDBIn struct {
	OrderDetailId *uint32
	Token         *string
	Label         string
	Trace         string
}

//+ NEW ITEM
type NewItemDBIn struct {
	NewItemDBInData OrderDetail
	Label           *string
	Trace           *string
}

//+ UPDATE ITEM
type UpdateItemDBIn struct {
	UpdateItemDBInData OrderDetail
	Label              *string
	Trace              *string
}
