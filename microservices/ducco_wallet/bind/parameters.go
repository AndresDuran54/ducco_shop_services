package bind

import (
	"ducco/core/router"
)

type ParameterItem struct {
	router.HeadersCredentials
	ParamId *string `json:"paramId"`
}
