package products

import (
	"net/http"

	"ducco/microservices/ducco_categories/bind"
	"ducco/microservices/ducco_categories/repository/categories"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (o Handler) ItemsCustomer(c echo.Context, itemsCustomer interface{}) error {
	//+ Obtenemos la data de la consulta
	data := itemsCustomer.(*bind.ItemsCustomer)

	//+ Instancia del repository de los productos
	categoriesData := categories.Data{}

	//+ Obtenemos los productos
	categoriesResultDB := categoriesData.ItemsDB(categories.ItemsDBIn{
		PagingSize:  data.PagingSize,
		PagingIndex: data.PagingIndex,
		FilterVals:  data.Filters,
		OrderVals:   data.Orders,
	})

	//+ Pipe
	categoriesResultDB.Data.Items = ItemsCustomer(categoriesResultDB.Data.Items.([]categories.Category))
	return c.JSON(http.StatusOK, categoriesResultDB)
}
