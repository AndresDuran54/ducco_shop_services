package orders

type Order struct {
	OrderId           *uint32 `gorm:"column:orderId;primaryKey;autoIncrement" json:"orderId"`
	CustomerId        *uint32 `gorm:"column:customerId" json:"customerId"`
	AddressCustomer   *string `gorm:"column:addressCustomer" json:"addressCustomer"`
	Amount            *uint32 `gorm:"column:amount" json:"amount"`
	PartialAmount     *uint32 `gorm:"column:partialAmount" json:"partialAmount"`
	DeliveryAmount    *uint32 `gorm:"column:deliveryAmount" json:"deliveryAmount"`
	Status            *uint8  `gorm:"column:status" json:"status"`
	PaymentMethodId   *uint32 `gorm:"column:paymentMethodId" json:"paymentMethodId"`
	DeliveryTimestamp *uint64 `gorm:"column:deliveryTimestamp" json:"deliveryTimestamp"`
	InsTimestamp      *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
	CancelTimestamp   *uint64 `gorm:"column:cancelTimestamp" json:"cancelTimestamp"`
}

type BuildWhere struct {
	OrderId *uint32 `db:"orderId" pattern:"="`
	Token   *string `db:"token" pattern:"="`
}

//+ ITEM
type ItemDBIn struct {
	OrderId *uint32
	Label   string
	Trace   string
}

//+ NEW ITEM
type NewItemDBIn struct {
	NewItemDBInData Order
	Label           *string
	Trace           *string
}

//+ UPDATE ITEM
type UpdateItemDBIn struct {
	UpdateItemDBInData Order
	Label              *string
	Trace              *string
}
