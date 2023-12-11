package config

import (
	"os"

	"github.com/joho/godotenv"
)

//+ Objeto que almacenara las variables de entorno
var Env = EnvInterface{}

func init() {
	//+ Cargamos las variables de entorno
	godotenv.Load()

	Env = EnvInterface{
		App: App{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		DB: DB{
			MySql: MySql{
				Host: os.Getenv("MYSQL_HOST"),
				Port: os.Getenv("MYSQL_PORT"),
				Name: os.Getenv("MYSQL_NAME"),
				User: os.Getenv("MYSQL_USER"),
				Pass: os.Getenv("MYSQL_PASS"),
			},
		},
	}

}
