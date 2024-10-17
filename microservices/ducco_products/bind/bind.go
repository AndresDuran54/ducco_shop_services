package bind

import "ducco/core/router"

// + CUSTOMER
type ItemsCustomer struct {
	router.HeadersCredentialsFiltersPagingOrder
}

type ItemCustomer struct {
	router.HeadersCredentials
	ProductId *uint32 `param:"id"`
}

// + INTERSERVICES
type ProductInterSVC struct {
	router.HeadersCredentialsInterSVC
	ProductId *uint32 `json:"productId"`
}
