package sessions

import (
	"ducco/core/database"
	"ducco/microservices/ducco_customers/lib"
)

const TableName = "Sessions"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"customerId": {
			Column:  "customerId",
			Pattern: database.BetweenPattern,
		},
	}
}

func Orders() map[string]database.Order {
	return map[string]database.Order{
		"customerId": {
			Column: "customerId",
		},
	}
}

type Data struct{}

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var sessions []Sessions

	//+ Obtenemos los productos
	sessionsResult := lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:       sessions,
		Last:        itemsDBIn.PagingSize,
		Offset:      itemsDBIn.PagingIndex,
		Filters:     Filters(),
		FiltersVals: itemsDBIn.FilterVals,
		Orders:      Orders(),
		OrdersVals:  itemsDBIn.OrderVals,
		TableName:   TableName,
	})

	return sessionsResult
}

func (o Data) ItemDB(itemDBIn ItemDBIn) database.ItemDBOut {
	//+ Construcción del where
	buildWhere := BuildWhere{
		CustomerId: itemDBIn.CustomerId,
		Token:      itemDBIn.Token,
	}

	//+ Obtenemos el registro del cliente
	sessionResult := lib.MYSQL.ItemDB(database.ItemDBIn{
		Item:       &Sessions{},
		TableName:  TableName,
		BuildWhere: buildWhere,
	})

	return sessionResult
}

func (o Data) NewItemDB(newItemDBIn NewItemDBIn) database.NewItemDBOut {
	return lib.MYSQL.NewItemDB(database.NewItemDBIn{
		Item:      &newItemDBIn.NewItemDBInData,
		TableName: TableName,
	})
}

func (o Data) UpdateItemsDB(updateItemsDBIn UpdateItemsDBIn) database.UpdateItemsDBOut {
	return lib.MYSQL.UpdateItemsDB(database.UpdateItemsDBIn{
		Data: updateItemsDBIn.Data,
		BuildWhere: BuildWhere{
			CustomerId: updateItemsDBIn.CustomerId,
		},
		TableName: TableName,
	})
}
