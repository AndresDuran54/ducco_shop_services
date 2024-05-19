package wallet

import (
	"ducco/microservices/ducco_customers/repository/customers"
	"ducco/microservices/ducco_customers/repository/sessions"
)

type ResponsePipe struct {
	Data interface{} `json:"data"`
}

type SessionsLoginPipe struct {
	Customer customers.Customers `json:"customer"`
	Session  sessions.Sessions   `json:"item"`
}

func SessionsLogin(customer customers.Customers, session sessions.Sessions) ResponsePipe {
	sessionsPipe := SessionsLoginPipe{
		Customer: customer,
		Session:  session,
	}

	return ResponsePipe{
		Data: sessionsPipe,
	}
}

type SessionTokenInfoPipe struct {
	Customer customers.Customers `json:"customer"`
	Session  sessions.Sessions   `json:"item"`
}

func SessionTokenInfo(customer customers.Customers, session sessions.Sessions) ResponsePipe {
	sessionsPipe := SessionTokenInfoPipe{
		Customer: customer,
		Session:  session,
	}

	return ResponsePipe{
		Data: sessionsPipe,
	}
}
