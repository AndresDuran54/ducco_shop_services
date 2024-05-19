package conflicts

type ErrorData struct {
	Data interface{} `json:"data"`
}

type BadRequest struct {
	Message string `json:"message"`
}

type ConflictData struct {
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

type InternalServerErrorData struct {
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

type UnauthorizedErrorData struct {
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

type ErrorConflicts struct {
	MessageId string
	Message   string
}

var ERR_INTERNAL_SERVER_ERROR = ErrorConflicts{
	MessageId: "ERR_INTERNAL_SERVER_ERROR",
	Message:   "ERR_INTERNAL_SERVER_ERROR",
}

var ERR_UNAUTHORIZED_ERROR = ErrorConflicts{
	MessageId: "ERR_UNAUTHORIZED_ERROR",
	Message:   "ERR_UNAUTHORIZED_ERROR",
}
