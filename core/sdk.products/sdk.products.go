package sdk_products

import (
	"ducco/core/conflicts"
	"ducco/core/network"
	"ducco/core/utils"
	"ducco/microservices/ducco_products/repository/products"
	"fmt"
)

type SDKProducts struct {
	Host       string
	ApiKey     string
	Routes     map[string]map[string]string
	HttpClient network.HttpClient
	UtilsInt   utils.UtilsInt
}

func NewSDKProducts(host string, apiKey string) SDKProducts {
	return SDKProducts{
		Host:       host,
		ApiKey:     apiKey,
		HttpClient: network.HttpClient{},
		UtilsInt:   utils.UtilsInt{},
		Routes: map[string]map[string]string{
			"products": {
				"search-item": "/v1/products/search-item/interservices",
			},
		},
	}
}

//+ Products
func (o *SDKProducts) ProductsSearchItem(data ProductsSearchItemDataIn) ProductsSearchItemDataOut {

	//+ Realizamos la solicitud HTTP
	response := o.HttpClient.Call(network.HttpClientRequest{
		Method: "POST",
		Url:    o.Host + o.Routes["products"]["search-item"],
		Headers: map[string]string{
			"api_key": o.ApiKey,
		},
		Data: data,
	})

	fmt.Println("response.Data")
	fmt.Println(response.Data)
	fmt.Println("response.Data")

	if !response.Success {
		conflicts.InternalServerError(conflicts.InternalServerErrorData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   conflicts.ERR_INTERNAL_SERVER_ERROR.Message,
		})
	}

	//+ Instancia del utilitario del mapa
	utilsMap := utils.UtilsMap{}

	//+ Objeto para almacenar el products
	var product products.Product

	//+ Parseamos el customer
	productsMap := response.Data["data"].(map[string]interface{})["item"]

	err := utilsMap.InterfaceToStruct(productsMap, &product)

	if err != nil {
		conflicts.InternalServerError(conflicts.InternalServerErrorData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   err.Error(),
		})
	}

	return ProductsSearchItemDataOut{
		Success: response.Success,
		Product: product,
	}
}
