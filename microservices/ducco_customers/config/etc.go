package config

//+ Objeto que almacenara las variables de entorno
var Etc = EtcInterface{}

func init() {
	Etc = EtcInterface{
		Sessions: Sessions{
			SessionsStatus: SessionsStatus{
				Active:   1,
				Inactive: 0,
			},
		},
	}

}
