package routes

import (
	"ducco/microservices/ducco_wallet/bind"
	c "ducco/microservices/ducco_wallet/config"
	handler "ducco/microservices/ducco_wallet/controller/payment_methods"
	"ducco/microservices/ducco_wallet/guards"
	"ducco/microservices/ducco_wallet/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesPaymentMethods(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "payment-methods"

	//+ Handler
	handler := handler.Handler{}

	//+ Crear una nueva orden
	e.GET(prefix, func(c echo.Context) error {
		return guards.Request(guards.RequestIn[bind.PaymentMethodsItems]{
			RequestDataIn: guards.RequestDataIn{
				C: c,
			},
			RequestData: &bind.PaymentMethodsItems{},
			CheckGuard:  true,
			GuardFunc:   guards.CheckCustomerSession,
			BindFunc:    lib.Bind{}.Bind,
			HandlerFunc: handler.PaymentMethodsItems,
		})
	})
}
