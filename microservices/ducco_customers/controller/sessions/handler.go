package customers

import (
	"net/http"

	"ducco/core/conflicts"
	"ducco/core/utils"
	"ducco/microservices/ducco_customers/bind"
	"ducco/microservices/ducco_customers/config"
	"ducco/microservices/ducco_customers/repository/customers"
	"ducco/microservices/ducco_customers/repository/sessions"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (o Handler) SessionsLogin(c echo.Context, sessionLoginData interface{}) error {

	//+ Obtenemos la data de la consulta
	data := sessionLoginData.(*bind.SessionsLogin)

	//+ Instancia del repository de los sessions
	sessionsData := sessions.Data{}

	//+ Instancia del repository de los customers
	customersData := customers.Data{}

	//+ Obtenemos el cliente por el email
	customerResult := customersData.ItemDB(customers.ItemDBIn{
		Email: data.Email,
	})

	if !customerResult.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_CUSTOMER_EMAIL_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_CUSTOMER_EMAIL_NOT_FOUND.Message,
		})
	}

	//+ Registro del cliente
	customer := customerResult.Data.Item.(*customers.Customers)

	//+ Comprobamos las contraseñas
	err := utils.UtilsCrypto{}.CompareHashAndPassword(*customer.Password, *data.Password)

	if err != nil {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_SESSIONS_PASSWORDS_NOT_MATCH.MessageId,
			Message:   err.Error(),
		})
	}

	//+ Obtenemos las marcas del tiempo actual
	currentTime := utils.UtilDate{}.CurrentTimeUTC()

	//TODO: OBTENER FECHA DE EXPIRACIÓN DE BASE DE DATOS

	//+ Agregamos la fecha de expiración
	expTime := utils.UtilDate{}.AddHours(currentTime.Original, 72)

	//+ Generamos el token
	token, err := utils.UtilsJWT{}.GenerateJWTToken(config.Env.JWT.SecretKey, expTime.TimeStamp, *customer.CustomerId)

	if err != nil {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   err.Error(),
		})
	}

	//+ Finalizamos todos los tokens
	sessionsData.UpdateItemsDB(sessions.UpdateItemsDBIn{
		CustomerId: customer.CustomerId,
		Data: sessions.Sessions{
			Status: &config.Etc.Sessions.SessionsStatus.Inactive,
		},
	})

	//+ Generamos el nuevo token
	sessionResult := sessionsData.NewItemDB(sessions.NewItemDBIn{
		NewItemDBInData: sessions.Sessions{
			CustomerId:   customer.CustomerId,
			Status:       &config.Etc.Sessions.SessionsStatus.Active,
			Token:        &token,
			ExpTimestamp: &expTime.TimeStamp,
			InsTimestamp: &currentTime.TimeStamp,
		},
	})

	//+ Registro de la sesión
	session := sessionResult.Data.Item.(*sessions.Sessions)

	//+ Pipe
	return c.JSON(http.StatusOK, SessionsLogin(
		*customer,
		*session,
	))
}

func (o Handler) SessionsTokenInfo(c echo.Context, sessionTokenInfoData interface{}) error {

	//+ Obtenemos la data de la consulta
	data := sessionTokenInfoData.(*bind.SessionsTokenInfo)

	//+ Verificamos si la sesión existe
	sessionsData := sessions.Data{}

	//+ Instancia del repository de los customers
	customersData := customers.Data{}

	//+ Obtenemos el registro de la sesión
	sessionResult := sessionsData.ItemDB(sessions.ItemDBIn{
		Token: data.Token,
	})

	if !sessionResult.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_SESSIONS_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_SESSIONS_NOT_FOUND.Message,
		})
	}

	//+ Registro de la sesión
	session := sessionResult.Data.Item.(*sessions.Sessions)

	//+ Verificamos si la sesión este activa
	if *session.Status == config.Etc.Sessions.SessionsStatus.Inactive {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_SESSIONS_NOT_ACTIVE.MessageId,
			Message:   conflicts.ERR_SESSIONS_NOT_ACTIVE.Message,
		})
	}

	//+ Obtenemos el cliente por el email
	customerResult := customersData.ItemDB(customers.ItemDBIn{
		CustomerId: session.CustomerId,
	})

	if !customerResult.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_CUSTOMER_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_CUSTOMER_NOT_FOUND.Message,
		})
	}

	//+ Registro del cliente
	customer := customerResult.Data.Item.(*customers.Customers)

	//+ Pipe
	return c.JSON(http.StatusOK, SessionTokenInfo(
		*customer,
		*session,
	))
}

//+ INTERSERVICES
func (o Handler) SessionsCustomerValidateInterSVC(c echo.Context, sessionValidateData interface{}) error {

	//+ Obtenemos la data de la consulta
	data := sessionValidateData.(*bind.SessionsCustomerValidateInterSVC)

	//+ Verificamos si la sesión existe
	sessionsData := sessions.Data{}

	//+ Instancia del repository de los customers
	customersData := customers.Data{}

	//+ Obtenemos el registro de la sesión
	sessionResult := sessionsData.ItemDB(sessions.ItemDBIn{
		Token: data.Token,
	})

	if !sessionResult.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_SESSIONS_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_SESSIONS_NOT_FOUND.Message,
		})
	}

	//+ Registro de la sesión
	session := sessionResult.Data.Item.(*sessions.Sessions)

	//+ Verificamos si la sesión este activa
	if *session.Status == config.Etc.Sessions.SessionsStatus.Inactive {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_SESSIONS_NOT_ACTIVE.MessageId,
			Message:   conflicts.ERR_SESSIONS_NOT_ACTIVE.Message,
		})
	}

	//+ Obtenemos el cliente por el customerId
	customerResult := customersData.ItemDB(customers.ItemDBIn{
		CustomerId: session.CustomerId,
	})

	if !customerResult.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_CUSTOMER_NOT_FOUND.MessageId,
			Message:   conflicts.ERR_CUSTOMER_NOT_FOUND.Message,
		})
	}

	//+ Registro del cliente
	customer := customerResult.Data.Item.(*customers.Customers)

	//+ Pipe
	return c.JSON(http.StatusOK, SessionCustomerValidateInterSVC(
		*customer,
		*session,
	))
}
