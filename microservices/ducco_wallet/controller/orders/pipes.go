package wallet

import (
	sdk_products "ducco/core/sdk.products"
	"ducco/microservices/ducco_wallet/repository/orders"
	"ducco/microservices/ducco_wallet/repository/orders_details"
	"encoding/json"
)

//+ GENERAL
type OrderAddressCustomer struct {
	Address        *string `json:"address"`
	District       *string `json:"district"`
	FloorApartment *string `json:"floorApartment"`
	Reference      *string `json:"reference"`
}

//+ NEW ORDER PIPE
type NewOrderPipe struct {
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

func OrdersNewItem(order orders.Order) NewOrderPipe {
	orderPipe := NewOrderPipe{}

	addressCustomer := &OrderAddressCustomer{}

	json.Unmarshal([]byte(*order.AddressCustomer), addressCustomer)

	orderPipe = NewOrderPipe{
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

//+ GET ORDER PIPE
type GetOrderDetailProductPipe struct {
	ProductId      *uint32 `gorm:"column:productId" json:"productId"`
	CardTitleFO    *string `gorm:"column:cardTitleFO" json:"cardTitleFO"`
	CardSubTitleFO *string `gorm:"column:cardSubTitleFO" json:"cardSubTitleFO"`
	CardImgUrlFO   *string `gorm:"column:cardImgUrlFO" json:"cardImgUrlFO"`
	InventoryPrice *uint32 `gorm:"column:inventoryPrice" json:"inventoryPrice"`
	Quantity       *uint32 `gorm:"column:inventorySalesQuantity" json:"inventorySalesQuantity"`
}

type GetOrderPipe struct {
	OrderId           *uint32                      `json:"orderId"`
	CustomerId        *uint32                      `json:"customerId"`
	AddressCustomer   *OrderAddressCustomer        `json:"addressCustomer"`
	Amount            *uint32                      `json:"amount"`
	PartialAmount     *uint32                      `json:"partialAmount"`
	DeliveryAmount    *uint32                      `json:"deliveryAmount"`
	Status            *uint8                       `json:"status"`
	PaymentMethodId   *uint32                      `json:"paymentMethodId"`
	DeliveryTimestamp *uint64                      `json:"deliveryTimestamp"`
	InsTimestamp      *uint64                      `json:"insTimestamp"`
	CancelTimestamp   *uint64                      `json:"cancelTimestamp"`
	OrderDetails      *[]GetOrderDetailProductPipe `json:"orderDetails"`
}

func GetOrderProduct(order orders.Order, orderDetails []orders_details.OrderDetail, productsMap map[uint32]sdk_products.Product) GetOrderPipe {
	//+ Dirección de la orden
	addressCustomer := &OrderAddressCustomer{}
	json.Unmarshal([]byte(*order.AddressCustomer), addressCustomer)

	//+ Detalles de la orden
	orderDetailsPipe := []GetOrderDetailProductPipe{}
	for _, od := range orderDetails {
		//+ Registro del producto
		product := productsMap[*od.ProductId]

		orderDetailsPipe = append(orderDetailsPipe, GetOrderDetailProductPipe{
			ProductId:      product.ProductId,
			CardTitleFO:    product.CardTitleFO,
			CardSubTitleFO: product.CardSubTitleFO,
			CardImgUrlFO:   product.CardImgUrlFO,
			InventoryPrice: product.InventoryPrice,
			Quantity:       od.Quantity,
		})
	}

	return GetOrderPipe{
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
		OrderDetails:      &orderDetailsPipe,
	}
}
