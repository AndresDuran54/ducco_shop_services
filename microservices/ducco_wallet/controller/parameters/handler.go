package wallet

import (
	"ducco/core/conflicts"
	"ducco/microservices/ducco_wallet/bind"
	"ducco/microservices/ducco_wallet/guards"
	"ducco/microservices/ducco_wallet/repository/parameters"
	"fmt"
	"net/http"
)

type Handler struct{}

func (o Handler) ParameterItemDB(c guards.RequestDataIn, parameterItemData interface{}) error {

	//+ Obtenemos la data de la consulta
	data := parameterItemData.(*bind.ParameterItem)

	//+ Instancia del repository para los parámetros
	parametersData := parameters.Data{}

	//+ Obtenemos el registro del parámetro
	parameterResult := parametersData.ItemDB(parameters.ItemDBIn{
		ParamId: data.ParamId,
	})

	if !parameterResult.Data.ItemFound {
		fmt.Println("AAA")
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_PARAMETER_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_PARAMETER_NOT_FOUND.Message,
		})
	}

	//+ Registro del parámetro
	parameter := parameterResult.Data.Item.(*parameters.Parameter)

	//+ Pipe
	parameterResult.Data.Item = ParameterItem(*parameter)
	return c.C.JSON(http.StatusOK, parameterResult)
}
