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

func (o Handler) ItemsCustomer(c echo.Context, itemsCustomer interface{}) error {
	//+ Obtenemos la data de la consulta
	data := itemsCustomer.(*bind.ItemsCustomer)

	//+ Instancia del repository de los productos
	customersData := customers.Data{}

	//+ Obtenemos los productos
	customersResultDB := customersData.ItemsDB(customers.ItemsDBIn{
		PagingSize:  data.PagingSize,
		PagingIndex: data.PagingIndex,
		FilterVals:  data.Filters,
		OrderVals:   data.Orders,
	})

	//+ Pipe
	return c.JSON(http.StatusOK, customersResultDB)
}

func (o Handler) CustomersNew(c echo.Context, itemsCustomer interface{}) error {

	//+ Obtenemos la data de la consulta
	data := itemsCustomer.(*bind.CustomersNew)

	//+ Instancia del repository de los productos
	customersData := customers.Data{}

	//+ Instancia del repository de los sessions
	sessionsData := sessions.Data{}

	//+ Verificamos que no exista un cliente registrado con el email
	customerResultDB := customersData.ItemDB(customers.ItemDBIn{
		Email: data.Email,
	})

	if customerResultDB.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_CUSTOMER_EMAIL_ALREADY_EXISTS.MessageId,
			Message:   conflicts.ERR_CUSTOMER_EMAIL_ALREADY_EXISTS.Message,
		})
	}

	//+ Verificamos que no exista un cliente registrado con la misma identificación
	customerResultDB = customersData.ItemDB(customers.ItemDBIn{
		Identification: data.Identification,
	})

	if customerResultDB.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_CUSTOMER_IDENTIFICATION_ALREADY_EXISTS.MessageId,
			Message:   conflicts.ERR_CUSTOMER_IDENTIFICATION_ALREADY_EXISTS.Message,
		})
	}

	//+ Verificamos que no exista un cliente registrado con el mismo número celular
	customerResultDB = customersData.ItemDB(customers.ItemDBIn{
		Phone: data.Phone,
	})

	if customerResultDB.Data.ItemFound {
		conflicts.Conflict(conflicts.ConflictData{
			MessageId: conflicts.ERR_CUSTOMER_PHONE_ALREADY_EXISTS.MessageId,
			Message:   conflicts.ERR_CUSTOMER_PHONE_ALREADY_EXISTS.Message,
		})
	}

	//+ Generamos el hash de la contraseña
	password, err := utils.UtilsCrypto{}.GenerateFromPassword(*data.Password)

	if err != nil {
		conflicts.InternalServerError(conflicts.InternalServerErrorData{
			MessageId: conflicts.ERR_INTERNAL_SERVER_ERROR.MessageId,
			Message:   err.Error(),
		})
	}

	//+ Obtenemos las marcas del tiempo actual
	currentTime := utils.UtilDate{}.CurrentTimeUTC()

	//+ Registramos el nuevo cliente
	customersResultDB := customersData.NewItemDB(customers.NewItemDBIn{
		NewItemDBInData: customers.Customers{
			FirstName:         data.FirstName,
			LastName:          data.LastName,
			IdentId:           data.IdentId,
			Identification:    data.Identification,
			Email:             data.Email,
			Password:          &password,
			PhoneNumber:       data.Phone,
			BirthdayTimestamp: data.BirthdayTimestamp,
			InsTimestamp:      &currentTime.TimeStamp,
		},
	})

	//+ Registro del customer
	customer := customersResultDB.Data.Item.(*customers.Customers)

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

	//+ Generamos el nuevo token
	sessionResult := sessionsData.NewItemDB(sessions.NewItemDBIn{
		NewItemDBInData: sessions.Sessions{
			CustomerId:   customer.CustomerId,
			Token:        &token,
			ExpTimestamp: &expTime.TimeStamp,
			InsTimestamp: &currentTime.TimeStamp,
		},
	})

	//+ Registro de la sesión
	session := sessionResult.Data.Item.(*sessions.Sessions)

	//+ Pipe
	return c.JSON(http.StatusOK, CustomerNewItem(*customer, *session))
}

func (o Handler) CustomersSearchItemInterSVC(c echo.Context, customerSearchItemData interface{}) error {
	//+ Obtenemos la data de la consulta
	data := customerSearchItemData.(*bind.CustomersSearchItemInterSVC)

	//+ Instancia del repository de los customers
	customersData := customers.Data{}

	//+ Obtenemos el registro del cliente
	customerResultDB := customersData.ItemDB(customers.ItemDBIn{
		CustomerId: data.CustomerId,
	})

	//+ Pipe
	return c.JSON(http.StatusOK, customerResultDB)
}
