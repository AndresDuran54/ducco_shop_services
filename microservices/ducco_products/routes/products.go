package routes

import (
	"ducco/microservices/ducco_products/bind"
	c "ducco/microservices/ducco_products/config"
	handler "ducco/microservices/ducco_products/controller/products"
	"ducco/microservices/ducco_products/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesProducts(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "products"

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

	//+ INTERSERVICES
	e.POST(prefix+"/search-item/interservices", func(c echo.Context) error {
		return Request(RequestIn[bind.ProductInterSVC]{
			c:           c,
			requestData: &bind.ProductInterSVC{},
			bindFunc:    lib.Bind{}.Bind,
			handlerFunc: handler.ProductInterSVC,
		})
	})

}
