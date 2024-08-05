package wallet

import (
	"net/http"

	"ducco/core/conflicts"
	sdk_products "ducco/core/sdk.products"
	"ducco/core/utils"
	"ducco/microservices/ducco_wallet/bind"
	"ducco/microservices/ducco_wallet/config"
	"ducco/microservices/ducco_wallet/guards"
	"ducco/microservices/ducco_wallet/lib"
	"ducco/microservices/ducco_wallet/repository/orders"
	"ducco/microservices/ducco_wallet/repository/orders_details"
)

type Handler struct{}

func (o Handler) OrdersNewItem(c guards.RequestDataIn, ordersNewItemData interface{}) error {

	//+ Obtenemos la data de la consulta
	data := ordersNewItemData.(*bind.OrdersNewItem)

	//+ Instancia del repository para los detalles
	ordersDetailsData := orders_details.Data{}

	//+ Instancia del repository para las ordenes
	ordersData := orders.Data{}

	//+ Obtenemos las marcas del tiempo actual
	currentTime := utils.UtilDate{}.CurrentTimeUTC()

	//+ Convertimos a string
	addressCustomerStr, err := utils.UtilsMap{}.InterfaceToMap(data.AddressCustomer)

	if err != nil {
		conflicts.InternalServerError(conflicts.InternalServerErrorData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   err.Error(),
		})
	}

	//+ Insertamos la nueva orden
	orderResult := ordersData.NewItemDB(orders.NewItemDBIn{
		NewItemDBInData: orders.Order{
			CustomerId:      c.SessionData.Customer.CustomerId,
			AddressCustomer: &addressCustomerStr,
			PaymentMethodId: data.PaymentMethodId,
			Status:          &config.Etc.Orders.OrdersStatus.Init,
			InsTimestamp:    &currentTime.TimeStamp,
		},
	})

	//+ Registro de la orden
	order := orderResult.Data.Item.(*orders.Order)

	//+ Calculamos el monto parcial de la orden iterando los detalles
	var orderPartialAmount uint32 = 0

	for _, d := range *data.OrdersDetails {
		//+ Obtenemos el registro del producto
		productResult := lib.SDKProducts.ProductsSearchItem(sdk_products.ProductsSearchItemDataIn{
			ProductId: *d.ProductId,
		})

		//+ Calculamos el monto total de la orden
		orderPartialAmount += *productResult.Product.InventoryPrice * *d.Quantity

		//+ Insertamos el detalle de la orden
		ordersDetailsData.NewItemDB(orders_details.NewItemDBIn{
			NewItemDBInData: orders_details.OrderDetail{
				OrderId:      order.OrderId,
				ProductId:    d.ProductId,
				Quantity:     d.Quantity,
				InsTimestamp: &currentTime.TimeStamp,
			},
		})
	}

	//+ Obtenemos el monto del delivery por defecto
	var deliveryAmount uint32 = 500

	//+ Calculamos el monto total de la orden
	var amount uint32 = deliveryAmount + orderPartialAmount

	//+ Actualizamos la orden de la solicitud
	orderUpdateResult := ordersData.UpdateItemDB(orders.UpdateItemDBIn{
		UpdateItemDBInData: orders.Order{
			OrderId:        order.OrderId,
			DeliveryAmount: &deliveryAmount,
			PartialAmount:  &orderPartialAmount,
			Amount:         &amount,
			Status:         &config.Etc.Orders.OrdersStatus.Pending,
		},
	})

	//+ Parseamos el registro de la orden
	order = orderUpdateResult.Data.Item.(*orders.Order)

	//+ Pipe
	orderUpdateResult.Data.Item = OrdersNewItem(*order)
	return c.C.JSON(http.StatusOK, orderUpdateResult)
}

func (o Handler) OrdersGetItem(c guards.RequestDataIn, OrdersGetItemData interface{}) error {
	//+ Obtenemos la data de la consulta
	data := OrdersGetItemData.(*bind.OrdersGetItem)

	//+ Instancia del repository para los detalles
	ordersDetailsData := orders_details.Data{}

	//+ Instancia del repository para las ordenes
	ordersData := orders.Data{}

	//+ Obtenemos el registro de la orden
	orderResult := ordersData.ItemDB(orders.ItemDBIn{
		OrderId: data.OrderId,
	})

	if !orderResult.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_ORDER_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_ORDER_NOT_FOUND.Message,
		})
	}

	//+ Registro de la orden
	order := orderResult.Data.Item.(*orders.Order)

	//+ Obtenemos los detalles de la orden
	orderDetailsResult := ordersDetailsData.ItemsDB(orders_details.ItemsDBIn{
		OrderId: data.OrderId,
	})

	//+ Registros de los detalles de las ordenes
	orderDetails := orderDetailsResult.Data.Items.([]orders_details.OrderDetail)

	//+ Mapa de productos. productId -> Product
	productsMap := map[uint32]sdk_products.Product{}

	//+ Obtenemos los registros de los productos
	for _, od := range orderDetails {
		//+ Obtenemos el registro del producto
		productResult := lib.SDKProducts.ProductsSearchItem(sdk_products.ProductsSearchItemDataIn{
			ProductId: *od.ProductId,
		})

		//+ Registro del producto
		product := productResult.Product

		//+ Agregamos el producto al mapa
		productsMap[*product.ProductId] = product
	}

	return c.C.JSON(http.StatusOK, GetOrderProduct(*order, orderDetails, productsMap))
}
