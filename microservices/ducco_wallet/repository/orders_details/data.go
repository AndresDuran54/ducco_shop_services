package orders_details

import (
	"ducco/core/database"
	"ducco/microservices/ducco_wallet/lib"
)

const TableName = "OrdersDetails"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"orderDetailId": {
			Column:  "orderDetailId",
			Pattern: database.EqualPattern,
		},
	}
}

func Orders() map[string]database.Order {
	return map[string]database.Order{
		"orderDetailId": {
			Column: "orderDetailId",
		},
	}
}

type Data struct{}

func (o Data) ItemDB(itemDBIn ItemDBIn) database.ItemDBOut {
	//+ Construcción del where
	buildWhere := BuildWhere{
		OrderDetailId: itemDBIn.OrderDetailId,
		Token:         itemDBIn.Token,
	}

	//+ Obtenemos el registro del cliente
	sessionResult := lib.MYSQL.ItemDB(database.ItemDBIn{
		Item:       &OrderDetail{},
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

func (o Data) UpdateItemDB(updateItemDBIn UpdateItemDBIn) database.UpdateItemDBOut {
	return lib.MYSQL.UpdateItemDB(database.UpdateItemDBIn{
		Data: &updateItemDBIn.UpdateItemDBInData,
		BuildWhere: BuildWhere{
			OrderDetailId: updateItemDBIn.UpdateItemDBInData.OrderDetailId,
		},
		TableName: TableName,
	})
}
