package routes

import (
	"ducco/microservices/ducco_wallet/bind"
	c "ducco/microservices/ducco_wallet/config"
	handler "ducco/microservices/ducco_wallet/controller/orders"
	"ducco/microservices/ducco_wallet/guards"
	"ducco/microservices/ducco_wallet/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesOrders(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "orders"

	//+ Handler
	handler := handler.Handler{}

	//+ Crear una nueva orden
	e.POST(prefix, func(c echo.Context) error {
		return guards.Request(guards.RequestIn[bind.OrdersNewItem]{
			RequestDataIn: guards.RequestDataIn{
				C: c,
			},
			RequestData: &bind.OrdersNewItem{},
			CheckGuard:  true,
			GuardFunc:   guards.CheckCustomerSession,
			BindFunc:    lib.Bind{}.Bind,
			HandlerFunc: handler.OrdersNewItem,
		})
	})
}
