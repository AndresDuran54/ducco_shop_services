package categories

type Category struct {
	CategoryId     *uint32 `gorm:"column:categoryId" json:"categoryId"`
	NameFO         *string `gorm:"column:nameFO" json:"nameFO"`
	CardTitleFO    *string `gorm:"column:cardTitleFO" json:"cardTitleFO"`
	CardSubTitleFO *string `gorm:"column:cardSubTitleFO" json:"cardSubTitleFO"`
	CardImgUrlFO   *string `gorm:"column:cardImgUrlFO" json:"cardImgUrlFO"`
	CardShowFO     *uint8  `gorm:"column:cardShowFO" json:"cardShowFO"`
	CardOrderFO    *uint8  `gorm:"column:cardOrderFO" json:"cardOrderFO"`
	InsTimestamp   *uint64 `gorm:"column:insTimestamp" json:"insTimestamp"`
}

type ItemsDBIn struct {
	CardOrderFO  *uint8
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
