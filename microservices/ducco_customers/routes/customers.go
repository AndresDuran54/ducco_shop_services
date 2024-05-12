package routes

import (
	"ducco/microservices/ducco_customers/bind"
	c "ducco/microservices/ducco_customers/config"
	handler "ducco/microservices/ducco_customers/controller/customers"
	"ducco/microservices/ducco_customers/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesCustomers(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "customers"

	//+ Handler
	handler := handler.Handler{}

	//+ Obtener registros de clientes
	e.GET(prefix, func(c echo.Context) error {
		return Request(RequestIn[bind.ItemsCustomer]{
			c:           c,
			requestData: &bind.ItemsCustomer{},
			bindFunc:    lib.Bind{}.Bind,
			handlerFunc: handler.ItemsCustomer,
		})
	})

	//+ Crear un nuevo cliente
	e.POST(prefix, func(c echo.Context) error {
		return Request(RequestIn[bind.NewCustomer]{
			c:           c,
			requestData: &bind.NewCustomer{},
			bindFunc:    lib.Bind{}.Bind,
			handlerFunc: handler.NewCustomer,
		})
	})
}
