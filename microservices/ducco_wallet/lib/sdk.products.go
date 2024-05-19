package lib

import (
	sdk_products "ducco/core/sdk.products"
	"ducco/microservices/ducco_wallet/config"
)

var SDKProducts sdk_products.SDKProducts

func init() {
	SDKProducts = sdk_products.NewSDKProducts(config.Env.SDKS.SDKSProducts.HOST, config.Env.App.ApiKey)
}
