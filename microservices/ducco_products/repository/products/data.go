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
		"inventoryStockGTE": {
			Column:  "inventoryStock",
			Pattern: database.GreaterThanPattern,
		},
		"categoryId": {
			Column:  "categoryId",
			Pattern: database.EqualPattern,
		},
		"brandIdIn": {
			Column:  "brandId",
			Pattern: database.InPattern,
		},
		"inventoryPriceBetween": {
			Column:  "inventoryPrice",
			Pattern: database.BetweenPattern,
		},
	}
}

func Orders() map[string]database.Order {
	return map[string]database.Order{
		"inventorySalesQuantity": {
			Column: "inventorySalesQuantity",
		},
		"inventoryStock": {
			Column: "inventoryStock",
		},
		"inventoryPrice": {
			Column: "inventoryPrice",
		},
	}
}

type Data struct{}

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var products []Product

	//+ Obtenemos los productos
	productsResult := lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:       products,
		Last:        itemsDBIn.PagingSize,
		Offset:      itemsDBIn.PagingIndex,
		Filters:     Filters(),
		FiltersVals: itemsDBIn.FilterVals,
		Orders:      Orders(),
		OrdersVals:  itemsDBIn.OrderVals,
		TableName:   TableName,
	})

	return productsResult
}
