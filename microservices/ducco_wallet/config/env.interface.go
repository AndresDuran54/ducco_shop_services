package config

type (
	EnvInterface struct {
		App  App
		DB   DB
		JWT  JWT
		SDKS SDKS
	}

	App struct {
		Host   string
		Port   string
		ApiKey string
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

	SDKS struct {
		SDKSCustomers SDKSCustomers
		SDKSProducts  SDKSProducts
	}

	SDKSCustomers struct {
		HOST string
	}

	SDKSProducts struct {
		HOST string
	}
)
