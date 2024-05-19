package payment_methods

type PaymentMethod struct {
	PaymentMethodId *uint32 `gorm:"column:paymentMethodId;primaryKey;autoIncrement" json:"paymentMethodId"`
	CardTitleFO     *string `gorm:"column:cardTitleFO" json:"cardTitleFO"`
	CardImgUrlFO    *string `gorm:"column:cardImgUrlFO" json:"cardImgUrlFO"`
	CardOrderFO     *uint32 `gorm:"column:cardOrderFO" json:"cardOrderFO"`
	InsTimestamp    *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type BuildWhere struct {
	PaymentMethodId *uint32 `db:"paymentMethodId" pattern:"="`
}

//+ ITEMS
type ItemsDBIn struct {
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
