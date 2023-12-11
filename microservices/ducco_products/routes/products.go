package routes

import (
	c "ducco/microservices/ducco_products/config"
	handler "ducco/microservices/ducco_products/controller/products"

	"github.com/labstack/echo/v4"
)

func loadRoutesProducts(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "products"

	//+ Handler
	handler := handler.Handler{}

	e.POST(prefix, handler.ItemsCustomer)
}
