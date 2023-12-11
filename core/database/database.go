package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MY_SQL = iota
	MONGO_DB
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
	Items      interface{}
	Offset     int
	Last       int
	TableName  string
	OrderBy    OrderBy
	BuildWhere interface{}
}

type ItemsDBOut struct {
	Items        interface{}
	ItemsCounter int
}

type OrderBy struct {
	Column string
	Desc   bool
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
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &MYSQL{
		gormDB: gormDB,
	}, nil
}
