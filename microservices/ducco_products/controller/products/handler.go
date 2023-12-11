package products

import (
	"net/http"

	"ducco/microservices/ducco_products/repository/products"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (o Handler) ItemsCustomer(c echo.Context) error {
	productsData := products.Data{}

	var StockGTE uint32 = 10

	return c.JSON(http.StatusOK, productsData.ItemsDB(products.ItemsDBIn{
		StockGTE: &StockGTE,
	}))
}
