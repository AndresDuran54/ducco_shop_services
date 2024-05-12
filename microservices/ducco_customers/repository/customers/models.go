package customers

type Customers struct {
	CustomerId        *uint32 `gorm:"column:customerId;primaryKey;autoIncrement" json:"customerId"`
	FirstName         *string `gorm:"column:firstName" json:"firstName"`
	LastName          *string `gorm:"column:lastName" json:"lastName"`
	IdentId           *string `gorm:"column:identId" json:"identId"`
	Identification    *string `gorm:"column:identification" json:"identification"`
	Email             *string `gorm:"column:email" json:"email"`
	Password          *string `gorm:"column:password" json:"password"`
	PhoneNumber       *string `gorm:"column:phoneNumber" json:"phoneNumber"`
	BirthdayTimestamp *uint64 `gorm:"column:birthdayTimestamp" json:"birthdayTimestamp"`
	InsTimestamp      *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type BuildWhere struct {
	CustomerId     *uint32 `db:"customerId" pattern:"="`
	Identification *string `db:"identification" pattern:"="`
	Email          *string `db:"email" pattern:"="`
	Phone          *string `db:"phoneNumber" pattern:"="`
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
	CustomerId     *uint32
	Identification *string
	Email          *string
	Phone          *string
	Label          string
	Trace          string
}

//+ NEW ITEM
type NewItemDBIn struct {
	NewItemDBInData Customers
	Label           *string
	Trace           *string
}
