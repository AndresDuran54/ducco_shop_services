package products

import (
	"ducco/core/database"
	"ducco/microservices/ducco_products/lib"
)

type Data struct{}

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var products []Product

	//+ Obtenemos los productos
	err, productsResult := lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:  products,
		Last:   10,
		Offset: 0,
		OrderBy: database.OrderBy{
			Column: "productId",
			Desc:   true,
		},
		TableName: "Products",
		BuildWhere: ProductsWhere{
			StockGTE: itemsDBIn.StockGTE,
			ProductIdIn: &[]uint32{
				1,
				2,
			},
		},
	})

	if err != nil {
		panic("AAAAAAAAAA")
	}

	return productsResult
}
