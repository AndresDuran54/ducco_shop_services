package sessions

type Sessions struct {
	SessionId    *uint32 `gorm:"column:sessionId;primaryKey;autoIncrement" json:"sessionId"`
	CustomerId   *uint32 `gorm:"column:customerId" json:"customerId"`
	Status       *uint8  `gorm:"column:status" json:"status"`
	Token        *string `gorm:"column:token" json:"token"`
	ExpTimestamp *uint64 `gorm:"column:expTimestamp" json:"expTimestamp"`
	InsTimestamp *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type BuildWhere struct {
	CustomerId *uint32 `db:"customerId" pattern:"="`
	Token      *string `db:"token" pattern:"="`
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

//+ ITEM
type ItemDBIn struct {
	CustomerId *uint32
	Token      *string
	Label      string
	Trace      string
}

//+ NEW ITEM
type NewItemDBIn struct {
	NewItemDBInData Sessions
	Label           *string
	Trace           *string
}

//+ UPDATE ITEM
type UpdateItemsDBIn struct {
	CustomerId *uint32
	Data       Sessions
	Label      *string
	Trace      *string
}
