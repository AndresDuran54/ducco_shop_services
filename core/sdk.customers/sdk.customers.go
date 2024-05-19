package sdk_customers

import (
	"ducco/core/conflicts"
	"ducco/core/network"
	"ducco/core/utils"
	"ducco/microservices/ducco_customers/repository/customers"
	"ducco/microservices/ducco_customers/repository/sessions"
)

type SDKCustomers struct {
	Host       string
	ApiKey     string
	Routes     map[string]map[string]string
	HttpClient network.HttpClient
	UtilsInt   utils.UtilsInt
}

func NewSDKCustomers(host string, apiKey string) SDKCustomers {
	return SDKCustomers{
		Host:       host,
		ApiKey:     apiKey,
		HttpClient: network.HttpClient{},
		UtilsInt:   utils.UtilsInt{},
		Routes: map[string]map[string]string{
			"customers": {
				"search-item": "/v1/customers/search-item/interservices",
			},
			"sessions": {
				"customer-validate": "/v1/sessions/customer-validate/interservices",
			},
		},
	}
}

//+ CUSTOMERS
func (o *SDKCustomers) CustomersSearchItem(data CustomersSearchItemDataIn) CustomersSearchItemDataOut {

	//+ Realizamos la solicitud HTTP
	response := o.HttpClient.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    o.Host + o.Routes["customers"]["search-item"],
		Headers: map[string]string{
			"api_key": o.ApiKey,
		},
		Data: data,
	})

	if !response.Success {
		return CustomersSearchItemDataOut{
			Success: response.Success,
		}
	}

	return CustomersSearchItemDataOut{
		Success: response.Success,
	}
}

//+ SESSIONS
func (o *SDKCustomers) SessionsCustomerValidate(data SessionsCustomerValidateDataIn) SessionsCustomerValidateDataOut {

	//+ Realizamos la solicitud HTTP
	response := o.HttpClient.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    o.Host + o.Routes["sessions"]["customer-validate"],
		Headers: map[string]string{
			"api_key": o.ApiKey,
		},
		Data: data,
	})

	if !response.Success {
		conflicts.UnauthorizedError(conflicts.UnauthorizedErrorData{
			MessageId: conflicts.ERR_UNAUTHORIZED_ERROR.MessageId,
			Message:   conflicts.ERR_UNAUTHORIZED_ERROR.Message,
		})
	}

	//+ Instancia del utilitario del mapa
	utilsMap := utils.UtilsMap{}

	//+ Objeto para almacenar el customer
	var customer customers.Customers

	//+ Objeto para almacenar la sesión
	var session sessions.Sessions

	//+ Parseamos el customer
	customerMap := response.Data["data"].(map[string]interface{})["customer"]

	err := utilsMap.InterfaceToStruct(customerMap, &customer)

	if err != nil {
		conflicts.InternalServerError(conflicts.InternalServerErrorData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   err.Error(),
		})
	}

	//+ Parseamos la sesión
	sessionMap := response.Data["data"].(map[string]interface{})["item"]

	err = utilsMap.InterfaceToStruct(sessionMap, &session)

	if err != nil {
		conflicts.InternalServerError(conflicts.InternalServerErrorData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   err.Error(),
		})
	}

	return SessionsCustomerValidateDataOut{
		Success:  response.Success,
		Customer: customer,
		Session:  session,
	}
}
