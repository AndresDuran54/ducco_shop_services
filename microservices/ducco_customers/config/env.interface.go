package config

type (
	EnvInterface struct {
		App App
		DB  DB
		JWT JWT
	}

	App struct {
		Host string
		Port string
	}

	DB struct {
		MySql MySql
	}

	MySql struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}

	JWT struct {
		SecretKey string
	}
)
