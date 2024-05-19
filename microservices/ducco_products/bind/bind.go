package bind

import "ducco/core/router"

//+ CUSTOMER
type ItemsCustomer struct {
	router.HeadersCredentialsFiltersPagingOrder
}

//+ INTERSERVICES
type ProductInterSVC struct {
	router.HeadersCredentialsInterSVC
	ProductId *uint32 `json:"productId"`
}
