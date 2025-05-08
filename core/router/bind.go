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
	CustomerId  *int    `header:"customer-id"`
	Token       *string `header:"token"`
	PagingIndex *int    `header:"paging-index"`
	PagingSize  *int    `header:"paging-size"`
	Filters     *string `header:"filters"`
}

type HeadersCredentialsFiltersPagingOrder struct {
	Token       *string `header:"token"`
	PagingIndex *int    `header:"paging-index"`
	PagingSize  *int    `header:"paging-size"`
	Filters     *string `header:"filters"`
	Orders      *string `header:"orders"`
}
