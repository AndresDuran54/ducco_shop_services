package router

type HeaderFilter struct {
	Filter *string `json:"filter"`
	Val    *string `json:"val"`
	Val2   *string `json:"val2"`
}

type HeadersCredentialsFiltersPaging struct {
	CustomerId  *int    `header:"customer_id"`
	Token       *string `header:"token"`
	PagingIndex *int    `header:"paging_index"`
	PagingSize  *int    `header:"paging_size"`
	Filters     *string `header:"filters"`
}
