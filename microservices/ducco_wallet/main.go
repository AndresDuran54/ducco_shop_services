package main

import (
	c "ducco/microservices/ducco_wallet/config"
	"fmt"

	r "ducco/microservices/ducco_wallet/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	//+ Nombre de la aplicación

	var echo *echo.Echo = echo.New()
	echo.HideBanner = false
	echo.HidePort = false

	var uri string = fmt.Sprintf("%s:%s", c.Env.App.Host, c.Env.App.Port)

	r.LoadRoutes(echo)
	echo.Start(uri)
}
