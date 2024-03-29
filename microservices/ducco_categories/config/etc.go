package config

type (
	EtcInterface struct {
		Categories Categories
	}

	Categories struct {
		OrderFO CategoriesOrderFO
	}

	CategoriesOrderFO struct {
		Yes uint8
		No  uint8
	}
)
