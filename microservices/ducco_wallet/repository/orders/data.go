package orders

import (
	"ducco/core/database"
	"ducco/microservices/ducco_wallet/lib"
)

const TableName = "Orders"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"customerId": {
			Column:  "customerId",
			Pattern: database.BetweenPattern,
		},
	}
}

func Orders_() map[string]database.Order {
	return map[string]database.Order{
		"customerId": {
			Column: "customerId",
		},
	}
}

type Data struct{}

func (o Data) ItemDB(itemDBIn ItemDBIn) database.ItemDBOut {
	//+ Construcción del where
	buildWhere := BuildWhere{
		OrderId: itemDBIn.OrderId,
	}

	//+ Obtenemos el registro del cliente
	orderResult := lib.MYSQL.ItemDB(database.ItemDBIn{
		Item:       &Order{},
		TableName:  TableName,
		BuildWhere: buildWhere,
	})

	return orderResult
}

func (o Data) NewItemDB(newItemDBIn NewItemDBIn) database.NewItemDBOut {
	return lib.MYSQL.NewItemDB(database.NewItemDBIn{
		Item:      &newItemDBIn.NewItemDBInData,
		TableName: TableName,
	})
}

func (o Data) UpdateItemDB(updateItemDBIn UpdateItemDBIn) database.ItemDBOut {

	//+ Actualizamos el registro
	lib.MYSQL.UpdateItemDB(database.UpdateItemDBIn{
		Data: &updateItemDBIn.UpdateItemDBInData,
		BuildWhere: BuildWhere{
			OrderId: updateItemDBIn.UpdateItemDBInData.OrderId,
		},
		TableName: TableName,
	})

	//+ Obtenemos el registro actualizado
	return o.ItemDB(ItemDBIn{
		OrderId: updateItemDBIn.UpdateItemDBInData.OrderId,
	})
}
