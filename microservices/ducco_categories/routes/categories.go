package routes

import (
	"ducco/core/conflicts"
	"ducco/microservices/ducco_categories/bind"
	c "ducco/microservices/ducco_categories/config"
	handler "ducco/microservices/ducco_categories/controller/categories"
	"ducco/microservices/ducco_categories/lib"
	"net/http"

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

func Request[T any](requestIn RequestIn[T]) error {

	defer func() error {
		if err := recover(); err != nil {
			return requestIn.c.JSON(http.StatusInternalServerError, conflicts.BadRequest{
				Message: err.(string),
			})
		}
		return nil
	}()

	err := requestIn.bindFunc(requestIn.c, requestIn.requestData)

	if err != nil {
		return requestIn.c.JSON(http.StatusBadRequest, conflicts.BadRequest{
			Message: err.Error(),
		})
	}

	return requestIn.handlerFunc(requestIn.c, requestIn.requestData)
}

type RequestIn[T any] struct {
	c           echo.Context
	requestData *T
	guardFunc   func(c echo.Context) error
	bindFunc    func(c echo.Context, bindModel interface{}) error
	handlerFunc func(c echo.Context, bindModel interface{}) error
}
