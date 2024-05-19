package wallet

import (
	"ducco/microservices/ducco_wallet/repository/orders"
	"encoding/json"
)

type OrderPipe struct {
	OrderId           *uint32               `json:"orderId"`
	CustomerId        *uint32               `json:"customerId"`
	AddressCustomer   *OrderAddressCustomer `json:"addressCustomer"`
	Amount            *uint32               `json:"amount"`
	PartialAmount     *uint32               `json:"partialAmount"`
	DeliveryAmount    *uint32               `json:"deliveryAmount"`
	Status            *uint8                `json:"status"`
	PaymentMethodId   *uint32               `json:"paymentMethodId"`
	DeliveryTimestamp *uint64               `json:"deliveryTimestamp"`
	InsTimestamp      *uint64               `json:"insTimestamp"`
	CancelTimestamp   *uint64               `json:"cancelTimestamp"`
}

type OrderAddressCustomer struct {
	Address        *string `json:"address"`
	District       *string `json:"district"`
	FloorApartment *string `json:"floorApartment"`
	Reference      *string `json:"reference"`
}

func OrdersNewItem(order orders.Order) OrderPipe {
	orderPipe := OrderPipe{}

	addressCustomer := &OrderAddressCustomer{}

	json.Unmarshal([]byte(*order.AddressCustomer), addressCustomer)

	orderPipe = OrderPipe{
		OrderId:           order.OrderId,
		CustomerId:        order.CustomerId,
		AddressCustomer:   addressCustomer,
		Amount:            order.Amount,
		PartialAmount:     order.PartialAmount,
		DeliveryAmount:    order.DeliveryAmount,
		Status:            order.Status,
		PaymentMethodId:   order.PaymentMethodId,
		DeliveryTimestamp: order.DeliveryTimestamp,
		InsTimestamp:      order.InsTimestamp,
		CancelTimestamp:   order.CancelTimestamp,
	}

	return orderPipe
}
