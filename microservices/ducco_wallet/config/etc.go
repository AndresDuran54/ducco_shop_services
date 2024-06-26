package config

//+ Objeto que almacenara las variables de entorno
var Etc = EtcInterface{}

func init() {
	Etc = EtcInterface{
		Orders: Orders{
			OrdersStatus: OrdersStatus{
				Init:      0,
				Success:   1,
				Delivered: 2,
				Error:     3,
				Pending:   4,
			},
		},
		Parameters: Parameters{
			ParametersTypes: ParametersTypes{
				Number: 0,
				String: 1,
				JSON:   2,
			},
		},
	}

}
