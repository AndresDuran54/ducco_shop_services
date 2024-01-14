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
	ItemDB(itemsDBIn ItemsDBIn) (error, ItemsDBOut)
	ItemsDB(itemsDBIn ItemsDBIn) (error, ItemsDBOut)
	NewItemDB() (error, interface{})
	UpdateItemDB() (error, interface{})
	BuildWhere(i interface{}) (string, []interface{})
}

type Database struct {
}

type ItemsDBIn struct {
	Items       interface{}
	Offset      *int
	Last        *int
	TableName   string
	OrderBy     *OrderBy
	BuildWhere  interface{}
	FiltersVals *string
	Filters     map[string]Filter
}

type ItemsDBOut struct {
	Data ItemDBDataOut `json:"data"`
}

type ItemDBDataOut struct {
	Items             interface{} `json:"items"`
	ItemsCounter      int         `json:"itemsCounter"`
	ItemsCounterTotal int         `json:"itemsCounterTotal"`
}

type OrderBy struct {
	Column string
	Desc   bool
}

type Filter struct {
	Column  string
	Pattern string
}

type FilterVals struct {
	Filter string `json:"filter"`
	Val    string `json:"val"`
	Val2   string `json:"val2"`
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
