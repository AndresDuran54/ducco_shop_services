package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	MY_SQL = iota
	MONGO_DB
)

//+ PATTERNS
const (
	EqualPattern              string = `=`
	NotEqualPattern           string = `!=`
	LikePattern               string = `LIKE`
	BetweenPattern            string = `BETWEEN`
	InPattern                 string = `IN`
	GreaterThanPattern        string = `>`
	GreaterThanOrEqualPattern string = `>=`
	LessThanPattern           string = `<`
	LessThanOrEqualPattern    string = `<=`
)

type IDatabase interface {
	ItemDB(itemDBIn ItemDBIn) ItemDBOut
	ItemsDB(itemsDBIn ItemsDBIn) ItemsDBOut
	NewItemDB(newItemDBIn NewItemDBIn) NewItemDBOut
	UpdateItemDB() (interface{}, error)
	UpdateItemsDB(updateItemsDBIn UpdateItemsDBIn) UpdateItemsDBOut
	BuildWhere(i interface{}) (string, []interface{})
}

type Database struct {
}

//+ +++++++++ ItemDB +++++++++

type ItemDBIn struct {
	Item       interface{}
	TableName  string
	BuildWhere interface{}
}

type ItemDBOut struct {
	Data ItemDBDataOut `json:"data"`
}

type ItemDBDataOut struct {
	ItemFound bool        `json:"itemFound"`
	Item      interface{} `json:"item"`
}

//+ +++++++++ ItemsDB +++++++++

type ItemsDBIn struct {
	Items       interface{}
	Offset      *int
	Last        *int
	TableName   string
	OrderBy     *OrderBy
	BuildWhere  interface{}
	FiltersVals *string
	Filters     map[string]Filter
	OrdersVals  *string
	Orders      map[string]Order
}

type ItemsDBOut struct {
	Data ItemsDBDataOut `json:"data"`
}

type ItemsDBDataOut struct {
	Items             interface{} `json:"items"`
	ItemsCounter      int         `json:"itemsCounter"`
	ItemsCounterTotal int         `json:"itemsCounterTotal"`
}

//+ +++++++++ NewItemDB +++++++++

type NewItemDBIn struct {
	Item      interface{}
	TableName string
}

type NewItemDBOut struct {
	Data NewItemDBDataOut `json:"data"`
}

type NewItemDBDataOut struct {
	Item interface{} `json:"item"`
}

//+ +++++++++ UpdateItemsDB +++++++++

type UpdateItemsDBIn struct {
	Data       interface{}
	BuildWhere interface{}
	TableName  string
}

type UpdateItemsDBOut struct {
	Data UpdateItemsDBDataOut `json:"data"`
}

type UpdateItemsDBDataOut struct {
	Item interface{} `json:"item"`
}

type OrderBy struct {
	Column string
	Desc   bool
}

type Filter struct {
	Column  string
	Pattern string
}

type Order struct {
	Column string
	Order  string
}

type FilterVals struct {
	Filter string `json:"filter"`
	Val    string `json:"val"`
	Val2   string `json:"val2"`
}

type OrderVals struct {
	Order string `json:"order"`
	Val   string `json:"val"`
}

type NewDatabaseIn struct {
	Type int
	User string
	Pass string
	Host string
	Port string
	Name string
}

func (o Database) NewDatabase(newDatabaseIn NewDatabaseIn) IDatabase {
	// Variables para almacenar el error
	var err error

	// Variable para almacenar la instancia nueva
	var database IDatabase

	// Escogemos el tipo de instancia
	switch newDatabaseIn.Type {
	case MY_SQL:
		database, err = newDatabaseMYSQL(newDatabaseIn)
	}

	// Comprobamos si hubo error
	if err != nil {
		panic("")
	}

	return database
}

func newDatabaseMYSQL(newDatabaseIn NewDatabaseIn) (IDatabase, error) {
	//+ Cadena de conexión
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", newDatabaseIn.User, newDatabaseIn.Pass, newDatabaseIn.Host, newDatabaseIn.Port, newDatabaseIn.Name)

	//+ Obtenemos la conexión
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return &MYSQL{
		gormDB: gormDB,
	}, nil
}
