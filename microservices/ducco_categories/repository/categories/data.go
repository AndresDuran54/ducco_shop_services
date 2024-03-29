package categories

import (
	"ducco/core/database"
	"ducco/microservices/ducco_categories/lib"
)

const TableName = "Categories"
const columnPK = "categoryId"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"insTimestampBetween": {
			Column:  "insTimestamp",
			Pattern: database.BetweenPattern,
		},
		"cardShowFO": {
			Column:  "cardShowFO",
			Pattern: database.EqualPattern,
		},
	}
}

func Orders() map[string]database.Order {
	return map[string]database.Order{
		"cardOrderFO": {
			Column: "cardOrderFO",
		},
	}
}

type Data struct{}

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var categories []Category

	//+ Obtenemos los productos
	categoriesResult := lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:       categories,
		Last:        itemsDBIn.PagingSize,
		Offset:      itemsDBIn.PagingIndex,
		Filters:     Filters(),
		FiltersVals: itemsDBIn.FilterVals,
		Orders:      Orders(),
		OrdersVals:  itemsDBIn.OrderVals,
		TableName:   TableName,
	})

	return categoriesResult
}
