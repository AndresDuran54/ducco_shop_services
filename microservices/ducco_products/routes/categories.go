package routes

import (
	"ducco/microservices/ducco_products/bind"
	c "ducco/microservices/ducco_products/config"
	handler "ducco/microservices/ducco_products/controller/categories"
	"ducco/microservices/ducco_products/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesCategories(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "categories"

	//+ Handler
	handler := handler.Handler{}

	e.GET(prefix, func(c echo.Context) error {
		return Request(RequestIn[bind.ItemsCustomer]{
			c:           c,
			requestData: &bind.ItemsCustomer{},
			bindFunc:    lib.Bind{}.Bind,
			handlerFunc: handler.ItemsCustomer,
		})
	})
}
