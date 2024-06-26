package database

import (
	"ducco/core/conflicts"
	"ducco/core/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MYSQL struct {
	gormDB *gorm.DB
}

func (o *MYSQL) ItemDB(itemDBIn ItemDBIn) ItemDBOut {

	//+ Definimos la tabla a consultar
	itemFindDB := o.gormDB.Table(itemDBIn.TableName).Unscoped()

	//+ Si se mando una configuración de construcción de WHERE
	if itemDBIn.BuildWhere != nil {
		//+ Obtenemos las sentencias condicionales
		whereStatement, whereStatementValues := o.BuildWhere(itemDBIn.BuildWhere)

		//+ Agregamos la sentencia where
		itemFindDB.Where(
			whereStatement,
			whereStatementValues...,
		)
	}

	//+ Obtenemos el registro
	itemDB := itemFindDB.First(itemDBIn.Item)

	return ItemDBOut{
		Data: ItemDBDataOut{
			ItemFound: itemDB.RowsAffected > 0,
			Item:      itemDBIn.Item,
		},
	}
}

func (o *MYSQL) ItemsDB(itemsDBIn ItemsDBIn) ItemsDBOut {

	//+ Definimos la tabla a consultar
	itemsFindDB := o.gormDB.Table(itemsDBIn.TableName).Unscoped()
	itemsTotalFindDB := o.gormDB.Table(itemsDBIn.TableName).Unscoped()

	//+ Si se mando una configuración de orden
	if itemsDBIn.OrderBy != nil {
		itemsFindDB = itemsFindDB.Order(clause.OrderByColumn{
			Column: clause.Column{Name: itemsDBIn.OrderBy.Column},
			Desc:   itemsDBIn.OrderBy.Desc,
		})
	}

	//+ Si se mando una configuración de paginado
	if itemsDBIn.Offset != nil && itemsDBIn.Last != nil {
		itemsFindDB = itemsFindDB.Offset((*itemsDBIn.Offset - 1) * *itemsDBIn.Last).Limit(*itemsDBIn.Last)
	}

	//+ Si se mando una configuración de construcción de WHERE
	if itemsDBIn.BuildWhere != nil {
		//+ Obtenemos las sentencias condicionales
		whereStatement, whereStatementValues := o.BuildWhere(itemsDBIn.BuildWhere)

		//+ Obtenemos los items con paginación
		itemsFindDB = itemsFindDB.Where(
			whereStatement,
			whereStatementValues...,
		)

		//+ Obtenemos los items sin paginación
		itemsTotalFindDB = itemsTotalFindDB.Where(
			whereStatement,
			whereStatementValues...,
		)
	}

	//+ Agregamos los filters
	if itemsDBIn.FiltersVals != nil && *itemsDBIn.FiltersVals != "" && len(itemsDBIn.Filters) > 0 {
		whereStatement, whereStatementValues := o.BuildFilters(itemsDBIn.Filters, *itemsDBIn.FiltersVals)

		//+ Obtenemos los items con paginación
		itemsFindDB = itemsFindDB.Where(
			whereStatement,
			whereStatementValues...,
		)

		//+ Obtenemos los items sin paginación
		itemsTotalFindDB = itemsTotalFindDB.Where(
			whereStatement,
			whereStatementValues...,
		)
	}

	//+ Agregamos los orders
	if itemsDBIn.OrdersVals != nil && *itemsDBIn.OrdersVals != "" && len(itemsDBIn.Orders) > 0 {
		var ordersVals = []OrderVals{}
		err := json.Unmarshal([]byte(*itemsDBIn.OrdersVals), &ordersVals)

		if err == nil {
			for _, orderVal := range ordersVals {
				itemsFindDB = itemsFindDB.Order(clause.OrderByColumn{
					Column: clause.Column{Name: orderVal.Order},
					Desc:   orderVal.Val == "desc",
				})
			}
		}
	}

	//+ Obtenemos el itemsCounterTotal
	itemsTotal := itemsDBIn.Items
	itemsTotalFindDB.Find(&itemsTotal)

	//+ Realizamos la consulta
	itemsFindDB.Find(&itemsDBIn.Items)

	//+ En caso de error
	if itemsFindDB.Error != nil {
		panic(itemsFindDB.Error.Error())
	}

	return ItemsDBOut{
		Data: ItemsDBDataOut{
			Items:             itemsDBIn.Items,
			ItemsCounter:      reflect.ValueOf(itemsDBIn.Items).Len(),
			ItemsCounterTotal: reflect.ValueOf(itemsTotal).Len(),
		},
	}
}

func (o *MYSQL) NewItemDB(newItemDBIn NewItemDBIn) NewItemDBOut {

	//+ Definimos la tabla a usar
	newItemDB := o.gormDB.Table(newItemDBIn.TableName).Unscoped()

	//+ Ingresamos el registro en base de datos
	newItemResult := newItemDB.Create(newItemDBIn.Item)

	if newItemResult.Error != nil {
		panic(conflicts.ErrorConflicts{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   newItemDB.Error.Error(),
		})
	}

	return NewItemDBOut{
		Data: NewItemDBDataOut{
			Item: newItemDBIn.Item,
		},
	}
}

func (o *MYSQL) UpdateItemDB(updateItemDBIn UpdateItemDBIn) UpdateItemDBOut {
	//+ Definimos la tabla a usar
	updateItemsDB := o.gormDB.Table(updateItemDBIn.TableName).Unscoped()

	//+ Si se mando una configuración de construcción de WHERE
	if updateItemDBIn.BuildWhere != nil {
		//+ Obtenemos las sentencias condicionales
		whereStatement, whereStatementValues := o.BuildWhere(updateItemDBIn.BuildWhere)

		//+ Agregamos la sentencia where
		updateItemsDB = updateItemsDB.Where(
			whereStatement,
			whereStatementValues...,
		)
	}

	//+ Ingresamos el registro en base de datos
	updateItemsResult := updateItemsDB.Updates(updateItemDBIn.Data)

	if updateItemsResult.Error != nil {
		panic(conflicts.ErrorConflicts{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   updateItemsDB.Error.Error(),
		})
	}

	return UpdateItemDBOut{
		Data: UpdateItemDBDataOut{
			Item: updateItemDBIn.Data,
		},
	}
}

func (o *MYSQL) UpdateItemsDB(updateItemsDBIn UpdateItemsDBIn) UpdateItemsDBOut {
	//+ Definimos la tabla a usar
	updateItemsDB := o.gormDB.Table(updateItemsDBIn.TableName).Unscoped()

	//+ Si se mando una configuración de construcción de WHERE
	if updateItemsDBIn.BuildWhere != nil {
		//+ Obtenemos las sentencias condicionales
		whereStatement, whereStatementValues := o.BuildWhere(updateItemsDBIn.BuildWhere)

		//+ Agregamos la sentencia where
		updateItemsDB = updateItemsDB.Where(
			whereStatement,
			whereStatementValues...,
		)
	}

	//+ Ingresamos el registro en base de datos
	updateItemsResult := updateItemsDB.Updates(updateItemsDBIn.Data)

	if updateItemsResult.Error != nil {
		panic(conflicts.ErrorConflicts{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   updateItemsDB.Error.Error(),
		})
	}

	return UpdateItemsDBOut{
		Data: UpdateItemsDBDataOut{
			Item: updateItemsDBIn.Data,
		},
	}
}

func (o *MYSQL) BuildWhere(i interface{}) (string, []interface{}) {
	var whereStatement []string
	var whereStatementValues []interface{}
	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)

	//+ Recorremos todos los campos del objeto
	for x := 0; x < refVal.NumField(); x++ {
		if refVal.Field(x).Kind() == reflect.Ptr && !refVal.Field(x).IsNil() {

			//+ Obtenemos el patrón
			pattern := refType.Field(x).Tag.Get("pattern")

			//+ Obtenemos el campo
			field := refVal.Field(x)

			//+ Obtenemos el tipo del campo
			elemKind := field.Elem().Kind()

			//+ Obtenemos el valor del tag db, que contiene el nombre de la columna
			col := refType.Field(x).Tag.Get("db")

			//+ Variable para almacenar el valor de la variable
			var val string

			//+ Actuamos según el pattern
			switch pattern {
			case "<>", "<", ">", "=":
				//+ Verificamos que sea un puntero string o un string, o un puntero a algún entero o un entero en sí
				if elemKind != reflect.String && !(elemKind >= reflect.Int && elemKind <= reflect.Uint64) {
					continue
				}

				//+ Obtenemos el valor del campo
				val = fmt.Sprintf("%v", field.Elem())

				//+ Si el campo está vació saltamos la siguiente iteración
				if val == "" {
					continue
				}

				//+ Agregamos la sentencia where
				whereStatement = append(whereStatement, fmt.Sprintf("%v %v ?", col, pattern))

				//+ Agregamos el valor de la sentencia
				whereStatementValues = append(whereStatementValues, val)
			case "IN":
				//+ Verificamos que sea un puntero a un slice o un slice
				if elemKind != reflect.Slice {
					continue
				}

				//+ Convertimos el reflect.Value a un *[]string
				sliceString, err := utils.ArrayUtils{}.ToSliceString(field)

				if err != nil {
					continue
				}

				//+ Si el campo está vació saltamos la siguiente iteración
				if len(*sliceString) == 0 {
					continue
				}

				//+ Crear un slice para almacenar los valores
				var val []string

				//+ Iterar sobre el map y agregar los valores al slice
				for _, value := range *sliceString {
					val = append(val, fmt.Sprintf("%v", value))
				}

				//+ Agregamos la sentencia where
				whereStatement = append(whereStatement, fmt.Sprintf("%v %v ?", col, pattern))

				//+ Agregamos el valor de la sentencia
				whereStatementValues = append(whereStatementValues, val)

			}
		}
	}

	return strings.Join(whereStatement, " AND "), whereStatementValues
}

func (o *MYSQL) BuildFilters(filters map[string]Filter, filtersVals string) (string, []interface{}) {
	var whereStatement []string
	var whereStatementValues []interface{}

	var filtersArray = []FilterVals{}

	err := json.Unmarshal([]byte(filtersVals), &filtersArray)

	if err != nil {
		return "", []interface{}{}
	}

	for _, filter := range filtersArray {

		//+ Si no existe el filtro saltamos la iteración
		if filters[filter.Filter].Column == "" || filters[filter.Filter].Pattern == "" {
			continue
		}

		//+ Variable para almacenar el valor de la variable
		var val, val2 string

		switch filters[filter.Filter].Pattern {
		case InPattern:
			//+ Obtenemos el valor del campo
			val = fmt.Sprintf("%v", filter.Val)

			//+ Si el campo está vació saltamos la siguiente iteración
			if val == "" {
				continue
			}

			//+ Agregamos la sentencia where
			whereStatement = append(whereStatement, fmt.Sprintf("%v %v (?)", filters[filter.Filter].Column, filters[filter.Filter].Pattern))

			//+ Agregamos el valor de la sentencia
			whereStatementValues = append(whereStatementValues, val)
		case EqualPattern, NotEqualPattern, LikePattern, GreaterThanPattern, GreaterThanOrEqualPattern, LessThanPattern, LessThanOrEqualPattern:

			//+ Obtenemos el valor del campo
			val = fmt.Sprintf("%v", filter.Val)

			//+ Si el campo está vació saltamos la siguiente iteración
			if val == "" {
				continue
			}

			//+ Agregamos la sentencia where
			whereStatement = append(whereStatement, fmt.Sprintf("%v %v ?", filters[filter.Filter].Column, filters[filter.Filter].Pattern))

			//+ Agregamos el valor de la sentencia
			whereStatementValues = append(whereStatementValues, val)
		case BetweenPattern:

			//+ Obtenemos el valor del campo
			val = fmt.Sprintf("%v", filter.Val)

			//+ Obtenemos el valor del campo
			val2 = fmt.Sprintf("%v", filter.Val2)

			//+ Si el campo está vació saltamos la siguiente iteración
			if val == "" {
				continue
			}

			//+ Si el campo está vació saltamos la siguiente iteración
			if val2 == "" {
				continue
			}

			//+ Agregamos la sentencia where
			whereStatement = append(whereStatement, fmt.Sprintf("%v %v ? AND ?", filters[filter.Filter].Column, filters[filter.Filter].Pattern))

			//+ Agregamos el valor de la sentencia
			whereStatementValues = append(whereStatementValues, val)
			whereStatementValues = append(whereStatementValues, val2)
		}
	}

	return strings.Join(whereStatement, " AND "), whereStatementValues
}
