package lib

import (
	"ducco/core/database"
	"ducco/microservices/ducco_products/config"
)

var MYSQL database.IDatabase

func init() {
	MYSQL = database.Database{}.NewDatabase(database.NewDatabaseIn{
		Type: database.MY_SQL,
		User: config.Env.DB.MySql.User,
		Pass: config.Env.DB.MySql.Pass,
		Host: config.Env.DB.MySql.Host,
		Port: config.Env.DB.MySql.Port,
		Name: config.Env.DB.MySql.Name,
	})
}
