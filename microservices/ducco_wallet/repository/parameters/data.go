package parameters

import (
	"ducco/core/database"
	"ducco/microservices/ducco_wallet/lib"
)

const TableName = "Parameters"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"paramId": {
			Column:  "paramId",
			Pattern: database.EqualPattern,
		},
	}
}

func Orders_() map[string]database.Order {
	return map[string]database.Order{
		"paramId": {
			Column: "paramId",
		},
	}
}

type Data struct{}

func (o Data) ItemDB(itemDBIn ItemDBIn) database.ItemDBOut {
	//+ Construcción del where
	buildWhere := BuildWhere{
		ParamId: itemDBIn.ParamId,
	}

	//+ Obtenemos el registro del cliente
	sessionResult := lib.MYSQL.ItemDB(database.ItemDBIn{
		Item:       &Parameter{},
		TableName:  TableName,
		BuildWhere: buildWhere,
	})

	return sessionResult
}
