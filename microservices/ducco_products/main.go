package main

import (
	c "ducco/microservices/ducco_products/config"
	"fmt"

	r "ducco/microservices/ducco_products/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	//+ Nombre de la aplicación

	var echo *echo.Echo = echo.New()
	echo.HideBanner = true
	echo.HidePort = true

	var uri string = fmt.Sprintf("%s:%s", c.Env.App.Host, c.Env.App.Port)

	r.LoadRoutes(echo)
	echo.Start(uri)
}
