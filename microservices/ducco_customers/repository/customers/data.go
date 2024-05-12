package customers

import (
	"ducco/core/database"
	"ducco/microservices/ducco_customers/lib"
)

const TableName = "Customers"

func Filters() map[string]database.Filter {
	return map[string]database.Filter{
		"identification": {
			Column:  "identification",
			Pattern: database.BetweenPattern,
		},
	}
}

func Orders() map[string]database.Order {
	return map[string]database.Order{
		"identification": {
			Column: "identification",
		},
	}
}

type Data struct{}

func (o Data) ItemsDB(itemsDBIn ItemsDBIn) database.ItemsDBOut {
	var customers []Customers

	//+ Obtenemos los productos
	customersResult := lib.MYSQL.ItemsDB(database.ItemsDBIn{
		Items:       customers,
		Last:        itemsDBIn.PagingSize,
		Offset:      itemsDBIn.PagingIndex,
		Filters:     Filters(),
		FiltersVals: itemsDBIn.FilterVals,
		Orders:      Orders(),
		OrdersVals:  itemsDBIn.OrderVals,
		TableName:   TableName,
	})

	return customersResult
}

func (o Data) ItemDB(itemDBIn ItemDBIn) database.ItemDBOut {
	//+ Construcción del where
	buildWhere := BuildWhere{
		CustomerId:     itemDBIn.CustomerId,
		Identification: itemDBIn.Identification,
		Email:          itemDBIn.Email,
		Phone:          itemDBIn.Phone,
	}

	//+ Obtenemos el registro del cliente
	customersResult := lib.MYSQL.ItemDB(database.ItemDBIn{
		Item:       &Customers{},
		TableName:  TableName,
		BuildWhere: buildWhere,
	})

	return customersResult
}

func (o Data) NewItemDB(newItemDBIn NewItemDBIn) database.NewItemDBOut {
	return lib.MYSQL.NewItemDB(database.NewItemDBIn{
		Item:      &newItemDBIn.NewItemDBInData,
		TableName: TableName,
	})
}
