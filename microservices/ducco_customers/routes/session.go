package routes

import (
	"ducco/microservices/ducco_customers/bind"
	c "ducco/microservices/ducco_customers/config"
	handler "ducco/microservices/ducco_customers/controller/sessions"
	"ducco/microservices/ducco_customers/lib"

	"github.com/labstack/echo/v4"
)

func loadRoutesSession(e *echo.Echo) {
	//+ Subject del paquete de rutas
	prefix := "/" + c.AppInfo.Version + "/" + "sessions"

	//+ Handler
	handler := handler.Handler{}

	//+ Crear un nuevo cliente
	e.POST(prefix+"/login", func(c echo.Context) error {
		return Request(RequestIn[bind.SessionsLogin]{
			c:           c,
			requestData: &bind.SessionsLogin{},
			bindFunc:    lib.Bind{}.Bind,
			handlerFunc: handler.SessionsLogin,
		})
	})

	//+ Obtener la info del cliente por su token
	e.POST(prefix+"/info", func(c echo.Context) error {
		return Request(RequestIn[bind.SessionsTokenInfo]{
			c:           c,
			requestData: &bind.SessionsTokenInfo{},
			bindFunc:    lib.Bind{}.Bind,
			handlerFunc: handler.SessionsTokenInfo,
		})
	})
}
