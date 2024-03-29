package config

//+ Objeto que almacenara las variables del micro
var Etc = EtcInterface{}

func init() {

	Etc = EtcInterface{
		Categories: Categories{
			OrderFO: CategoriesOrderFO{
				Yes: 1,
				No:  0,
			},
		},
	}

}
