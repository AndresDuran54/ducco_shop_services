package bind

import "ducco/core/router"

//+ CUSTOMER
type ItemsCustomer struct {
	router.HeadersCredentialsFiltersPaging
	Username  *string `json:"username"`
	Username2 *string `json:"username2" validate:"required"`
}
