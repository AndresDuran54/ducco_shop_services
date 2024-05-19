package router

type HeaderFilter struct {
	Filter *string `json:"filter"`
	Val    *string `json:"val"`
	Val2   *string `json:"val2"`
}

type HeadersCredentials struct {
	Token *string `header:"token"`
}

type HeadersCredentialsInterSVC struct {
	Token *string `header:"token"`
}

type HeadersCredentialsFiltersPaging struct {
	CustomerId  *int    `header:"customer_id"`
	Token       *string `header:"token"`
	PagingIndex *int    `header:"paging_index"`
	PagingSize  *int    `header:"paging_size"`
	Filters     *string `header:"filters"`
}

type HeadersCredentialsFiltersPagingOrder struct {
	Token       *string `header:"token"`
	PagingIndex *int    `header:"paging_index"`
	PagingSize  *int    `header:"paging_size"`
	Filters     *string `header:"filters"`
	Orders      *string `header:"orders"`
}
