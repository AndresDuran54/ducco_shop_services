package conflicts

func Conflict(conflictData ConflictData) {
	panic(conflictData)
}

func InternalServerError(internalServerErrorData InternalServerErrorData) {
	panic(internalServerErrorData)
}

func UnauthorizedError(unauthorizedErrorData UnauthorizedErrorData) {
	panic(unauthorizedErrorData)
}
