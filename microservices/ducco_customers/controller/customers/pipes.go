package customers

import (
	"ducco/microservices/ducco_customers/repository/customers"
	"ducco/microservices/ducco_customers/repository/sessions"
)

type ResponsePipe struct {
	Data interface{} `json:"data"`
}

type CustomerNewItemPipe struct {
	Customer customers.Customers `json:"item"`
	Session  sessions.Sessions   `json:"session"`
}

func CustomerNewItem(customer customers.Customers, session sessions.Sessions) ResponsePipe {
	sessionsPipe := CustomerNewItemPipe{
		Customer: customer,
		Session:  session,
	}

	return ResponsePipe{
		Data: sessionsPipe,
	}
}
