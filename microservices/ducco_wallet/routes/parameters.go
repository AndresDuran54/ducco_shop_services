package routes

import (
	"ducco/microservices/ducco_wallet/bind"
	c "ducco/microservices/ducco_wallet/config"
	handler "ducco/microservices/ducco_wallet/controller/parameters"
	"ducco/microservices/ducco_wallet/guards"
	"ducco/microservices/ducco_wallet/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesParameters(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "parameters"

	//+ Handler
	handler := handler.Handler{}

	//+ Obtener el registro de un parámetro
	e.GET(prefix+"/:id", func(c echo.Context) error {
		return guards.Request(guards.RequestIn[bind.ParameterItem]{
			RequestDataIn: guards.RequestDataIn{
				C: c,
			},
			RequestData: &bind.ParameterItem{},
			CheckGuard:  true,
			GuardFunc:   guards.CheckCustomerSession,
			BindFunc:    lib.Bind{}.Bind,
			HandlerFunc: handler.ParameterItemDB,
		})
	})
}
