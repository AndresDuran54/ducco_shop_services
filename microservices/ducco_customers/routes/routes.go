package routes

import (
	"ducco/core/conflicts"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoadRoutes(e *echo.Echo) {
	//+ Cargamos las rutas de los productos
	loadRoutesCustomers(e)
	loadRoutesSession(e)
}

func Request[T any](requestIn RequestIn[T]) error {

	defer func() error {
		if err := recover(); err != nil {
			return requestIn.c.JSON(http.StatusInternalServerError, conflicts.ErrorData{
				Data: err,
			})
		}
		return nil
	}()

	err := requestIn.bindFunc(requestIn.c, requestIn.requestData)

	if err != nil {
		return requestIn.c.JSON(http.StatusBadRequest, conflicts.ErrorData{
			Data: conflicts.BadRequest{
				Message: err.Error(),
			},
		})
	}

	if requestIn.useGuard {
		requestIn.guardFunc(requestIn.c)
	}

	return requestIn.handlerFunc(requestIn.c, requestIn.requestData)
}

type RequestIn[T any] struct {
	c           echo.Context
	useGuard    bool
	requestData *T
	guardFunc   func(c echo.Context) error
	bindFunc    func(c echo.Context, bindModel interface{}) error
	handlerFunc func(c echo.Context, bindModel interface{}) error
}
