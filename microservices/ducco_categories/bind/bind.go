package bind

import "ducco/core/router"

//+ CUSTOMER
type ItemsCustomer struct {
	router.HeadersCredentialsFiltersPagingOrder
}
