package payment_methods

import (
	"ducco/core/database"
	"ducco/microservices/ducco_wallet/lib"
)

const TableName = "PaymentMethods"

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

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var paymentMethods []PaymentMethod

	//+ Obtenemos el registro del cliente
	return lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:     paymentMethods,
		TableName: TableName,
	})
}
