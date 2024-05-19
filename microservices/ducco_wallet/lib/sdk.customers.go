package lib

import (
	sdk_customers "ducco/core/sdk.customers"
	"ducco/microservices/ducco_wallet/config"
)

var SDKCustomers sdk_customers.SDKCustomers

func init() {
	SDKCustomers = sdk_customers.NewSDKCustomers(config.Env.SDKS.SDKSCustomers.HOST, config.Env.App.ApiKey)
}
