package database

import (
	"ducco/core/utils"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MYSQL struct {
	gormDB *gorm.DB
}

func (o *MYSQL) ItemDB(itemsDBIn ItemsDBIn) (error, ItemsDBOut) {
	return nil, ItemsDBOut{}
}

func (o *MYSQL) ItemsDB(itemsDBIn ItemsDBIn) (error, ItemsDBOut) {

	var resultFind *gorm.DB

	if itemsDBIn.BuildWhere != nil {
		whereStatement, whereStatementValues := o.BuildWhere(itemsDBIn.BuildWhere)

		resultFind = o.gormDB.Table(itemsDBIn.TableName).Unscoped().Order(clause.OrderByColumn{
			Column: clause.Column{Name: itemsDBIn.OrderBy.Column},
			Desc:   itemsDBIn.OrderBy.Desc,
		}).Limit(itemsDBIn.Last).Offset(itemsDBIn.Offset).Where(
			whereStatement,
			whereStatementValues...,
		).Find(&itemsDBIn.Items)
	} else {
		resultFind = o.gormDB.Table(itemsDBIn.TableName).Unscoped().Order(clause.OrderByColumn{
			Column: clause.Column{Name: itemsDBIn.OrderBy.Column},
			Desc:   itemsDBIn.OrderBy.Desc,
		}).Limit(itemsDBIn.Last).Offset(itemsDBIn.Offset).Find(&itemsDBIn.Items)
	}

	if resultFind.Error != nil {
		return resultFind.Error, ItemsDBOut{}
	}

	return nil, ItemsDBOut{
		Items:        itemsDBIn.Items,
		ItemsCounter: reflect.ValueOf(itemsDBIn.Items).Len(),
	}
}

func (o *MYSQL) NewItemDB() (error, interface{}) {

	return nil, nil
}

func (o *MYSQL) UpdateItemDB() (error, interface{}) {

	return nil, nil
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
				fmt.Println("pattern " + pattern)
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
