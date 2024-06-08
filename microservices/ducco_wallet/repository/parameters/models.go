package parameters

type Parameter struct {
	ParamId      *string `gorm:"column:paramId" json:"paramId"`
	Value        *string `gorm:"column:value" json:"value"`
	Description  *string `gorm:"column:description" json:"description"`
	Type         *uint8  `gorm:"column:type" json:"type"`
	InsTimestamp *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type BuildWhere struct {
	ParamId *string `db:"paramId" pattern:"="`
}

//+ ITEM
type ItemDBIn struct {
	ParamId *string
	Label   string
	Trace   string
}
