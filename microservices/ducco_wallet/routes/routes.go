package routes

import (
	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo) {
	//+ Cargamos las rutas de los productos
	loadRoutesOrders(e)
	loadRoutesPaymentMethods(e)
	loadRoutesParameters(e)
}
