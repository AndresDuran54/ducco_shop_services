package bind

import "ducco/core/router"

//+ CUSTOMER
type ItemsCustomer struct {
	router.HeadersCredentialsFiltersPagingOrder
}

type NewCustomer struct {
	FirstName         *string `json:"firstName"`
	LastName          *string `json:"lastName"`
	IdentId           *string `json:"identId"`
	Identification    *string `json:"identification"`
	Email             *string `json:"email"`
	Password          *string `json:"password"`
	Phone             *string `json:"phoneNumber"`
	BirthdayTimestamp *uint64 `json:"birthdayTimestamp"`
}
