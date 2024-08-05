package bind

import (
	"ducco/core/router"
)

type OrdersNewItem struct {
	router.HeadersCredentials
	OrdersDetails   *[]OrdersNewItemDetail `json:"ordersDetails"`
	AddressCustomer *OrdersNewItemAddress  `json:"addressCustomer"`
	PaymentMethodId *uint32                `json:"paymentMethodId"`
}

type OrdersNewItemDetail struct {
	ProductId *uint32 `json:"productId"`
	Quantity  *uint32 `json:"quantity"`
}

type OrdersNewItemAddress struct {
	Address        *string `json:"address"`
	District       *string `json:"district"`
	FloorApartment *string `json:"floorApartment"`
	Reference      *string `json:"reference"`
}

type OrdersGetItem struct {
	router.HeadersCredentials
	OrderId *uint32 `param:"id"`
}
