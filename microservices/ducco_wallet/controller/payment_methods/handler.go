package wallet

import (
	"net/http"

	"ducco/microservices/ducco_wallet/bind"
	"ducco/microservices/ducco_wallet/guards"
	"ducco/microservices/ducco_wallet/repository/payment_methods"
)

type Handler struct{}

func (o Handler) PaymentMethodsItems(c guards.RequestDataIn, itemsCustomer interface{}) error {
	//+ Obtenemos la data de la consulta
	data := itemsCustomer.(*bind.PaymentMethodsItems)

	//+ Instancia del repository de los productos
	paymentMethodsData := payment_methods.Data{}

	//+ Obtenemos los productos
	paymentMethodsResultDB := paymentMethodsData.ItemsDB(payment_methods.ItemsDBIn{
		PagingSize:  data.PagingSize,
		PagingIndex: data.PagingIndex,
		FilterVals:  data.Filters,
		OrderVals:   data.Orders,
	})

	//+ Pipe
	return c.C.JSON(http.StatusOK, paymentMethodsResultDB)
}
