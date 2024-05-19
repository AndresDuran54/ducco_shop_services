package guards

import (
	"ducco/core/conflicts"
	sdk_customers "ducco/core/sdk.customers"
	"ducco/microservices/ducco_wallet/lib"
	"net/http"
	"reflect"
	"runtime"

	"github.com/labstack/echo/v4"
)

type RequestDataIn struct {
	C           echo.Context
	SessionData sdk_customers.SessionsCustomerValidateDataOut
}

type RequestIn[T any] struct {
	RequestDataIn RequestDataIn
	CheckGuard    bool
	RequestData   *T
	GuardFunc     func(token string) sdk_customers.SessionsCustomerValidateDataOut
	BindFunc      func(c echo.Context, bindModel interface{}) error
	HandlerFunc   func(c RequestDataIn, bindModel interface{}) error
}

func Request[T any](requestIn RequestIn[T]) error {

	defer func() error {
		if err := recover(); err != nil {
			//+ Error por defecto
			httpError := http.StatusInternalServerError

			//+ Error data
			var errorData any

			//+ Dependiendo del tipo escogemos el código de respuesta
			switch reflect.TypeOf(err).String() {
			case "conflicts.UnauthorizedErrorData":
				httpError = http.StatusUnauthorized
				errorData = err
			case "*runtime.TypeAssertionError":
				err_ := err.(*runtime.TypeAssertionError)
				errorData = conflicts.ErrorData{
					Data: conflicts.ErrorConflicts{
						MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
						Message:   err_.Error(),
					},
				}
			}

			return requestIn.RequestDataIn.C.JSON(httpError, conflicts.ErrorData{
				Data: errorData,
			})
		}
		return nil
	}()

	err := requestIn.BindFunc(requestIn.RequestDataIn.C, requestIn.RequestData)

	if err != nil {
		return requestIn.RequestDataIn.C.JSON(http.StatusBadRequest, conflicts.ErrorData{
			Data: conflicts.BadRequest{
				Message: err.Error(),
			},
		})
	}

	var sessionData sdk_customers.SessionsCustomerValidateDataOut
	if requestIn.CheckGuard {
		token := requestIn.RequestDataIn.C.Request().Header.Get("token")
		sessionData = requestIn.GuardFunc(token)
	}

	return requestIn.HandlerFunc(RequestDataIn{
		C:           requestIn.RequestDataIn.C,
		SessionData: sessionData,
	}, requestIn.RequestData)
}

func CheckCustomerSession(token string) sdk_customers.SessionsCustomerValidateDataOut {
	return lib.SDKCustomers.SessionsCustomerValidate(sdk_customers.SessionsCustomerValidateDataIn{
		Token: token,
	})
}
