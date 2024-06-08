package wallet

import (
	"ducco/core/utils"
	"ducco/microservices/ducco_wallet/config"
	"ducco/microservices/ducco_wallet/repository/parameters"
)

type ParameterPipe struct {
	ParamId      *string `json:"paramId"`
	Value        any     `json:"value"`
	Description  *string `json:"description"`
	Type         *uint8  `json:"type"`
	InsTimestamp *uint64 `json:"insTimestamp"`
}

func ParameterItem(parameter parameters.Parameter) ParameterPipe {
	//+ Pipe del parámetro
	parameterPipe := ParameterPipe{
		ParamId:      parameter.ParamId,
		Value:        parameter.Value,
		Description:  parameter.Description,
		Type:         parameter.Type,
		InsTimestamp: parameter.InsTimestamp,
	}

	//+ Dependiendo del tipo de dato lo parseamos
	switch *parameter.Type {
	case config.Etc.Parameters.ParametersTypes.Number:
		parameterPipe.Value = utils.UtilsInt{}.StringToInt64(*parameter.Value)
	}
	return parameterPipe
}
