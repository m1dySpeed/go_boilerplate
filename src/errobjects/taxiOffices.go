package errobjects

var (
	TOCreateFail = ErrorMessage{
		Message:     "Create taxi office error",
		Description: "Failed to create a new taxi office",
		Code:        400,
		CodeStr:     "E_CREATE_FAIL",
	}
	TOUpdateFail = ErrorMessage{
		Message:     "Update taxi office error",
		Description: "Failed to update taxi office",
		Code:        400,
		CodeStr:     "E_UPDATE_FAIL",
	}
	TODeleteFail = ErrorMessage{
		Message:     "Delete city error",
		Description: "Failed to delete a city",
		Code:        400,
		CodeStr:     "E_DELETE_FAIL",
	}
	TOWrongIdentifier = ErrorMessage{
		Message:     "Wrong taxi office Id",
		Description: "Please provide the taxi office identifier",
		Code:        400,
		CodeStr:     "E_WRONG_IDENTIFIER",
	}
	TORequestError = ErrorMessage{
		Message:     "Request error",
		Description: "Something wrong with your request",
		Code:        400,
		CodeStr:     "E_WRONG_REQUEST",
	}
	TONotFound = ErrorMessage{
		Message:     "Taxi office not found",
		Description: "Taxi office with provided id not found",
		Code:        404,
		CodeStr:     "E_WRONG_IDENTIFIER",
	}
)
