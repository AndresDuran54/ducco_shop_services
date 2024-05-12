package products

import (
	"fmt"
	"net/http"

	"ducco/microservices/ducco_products/bind"
	"ducco/microservices/ducco_products/repository/products"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (o Handler) ItemsCustomer(c echo.Context, itemsCustomer interface{}) error {
	//+ Obtenemos la data de la consulta
	data := itemsCustomer.(*bind.ItemsCustomer)

	//+ Instancia del repository de los productos
	productsData := products.Data{}

	//+ Obtenemos los productos
	productsResultDB := productsData.ItemsDB(products.ItemsDBIn{
		PagingSize:  data.PagingSize,
		PagingIndex: data.PagingIndex,
		FilterVals:  data.Filters,
		OrderVals:   data.Orders,
	})
	fmt.Println(*data.Orders)
	//+ Pipe
	productsResultDB.Data.Items = ItemsCustomer(productsResultDB.Data.Items.([]products.Product))
	return c.JSON(http.StatusOK, productsResultDB)
}
