package products

import (
	"ducco/core/database"
	"ducco/microservices/ducco_products/lib"
)

const TableName = "Products"
const columnPK = "productId"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"insTimestampBetween": {
			Column:  "insTimestamp",
			Pattern: database.BetweenPattern,
		},
		"priceGTE": {
			Column:  "price",
			Pattern: database.GreaterThanPattern,
		},
	}
}

type Data struct{}

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var products []Product

	//+ Obtenemos los productos
	err, productsResult := lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:       products,
		Last:        itemsDBIn.PagingSize,
		Offset:      itemsDBIn.PagingIndex,
		FiltersVals: itemsDBIn.FilterVals,
		Filters:     Filters(),
		TableName:   TableName,
		BuildWhere: ProductsWhere{
			StockGTE: itemsDBIn.StockGTE,
		},
		OrderBy: &database.OrderBy{
			Column: columnPK,
			Desc:   false,
		},
	})

	if err != nil {
		panic(err)
	}

	return productsResult
}
