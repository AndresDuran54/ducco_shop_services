package sdk_customers

import (
	"ducco/microservices/ducco_customers/repository/customers"
	"ducco/microservices/ducco_customers/repository/sessions"
)

//+ Customers
type CustomersSearchItemDataIn struct {
	CustomerId uint32
}

type CustomersSearchItemDataOut struct {
	Success    bool   `json:"success"`
	CustomerId string `json:"customerId"`
}

//+ Sessions
type SessionsValidateApiKeyDataIn struct {
	Token string
}

type SessionsValidateApiKeyDataOut struct {
	Success bool `json:"success"`
}

type SessionsCustomerValidateDataIn struct {
	Token string
}

type SessionsCustomerValidateDataOut struct {
	Success  bool                `json:"success"`
	Customer customers.Customers `json:"customer"`
	Session  sessions.Sessions   `json:"item"`
}
