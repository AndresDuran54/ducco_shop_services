package products

import (
	"net/http"

	"ducco/core/conflicts"
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

	//+ Pipe
	productsResultDB.Data.Items = ItemsCustomer(productsResultDB.Data.Items.([]products.Product))
	return c.JSON(http.StatusOK, productsResultDB)
}

//+ INTERSERVICES
func (o Handler) ProductInterSVC(c echo.Context, productData interface{}) error {
	//+ Obtenemos la data de la consulta
	data := productData.(*bind.ProductInterSVC)

	//+ Instancia del repository de los productos
	productsData := products.Data{}

	//+ Obtenemos el registro del producto
	productResultDB := productsData.ItemDB(products.ItemDBIn{
		ProductId: data.ProductId,
	})

	//+ Verificamos si existe o no el producto
	if !productResultDB.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_PRODUCT_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_PRODUCT_NOT_FOUND.Message,
		})
	}

	//+ Pipe
	return c.JSON(http.StatusOK, productResultDB)
}
