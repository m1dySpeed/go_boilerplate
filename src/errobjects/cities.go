package errobjects

var (
	CityCreateFail = ErrorMessage{
		Message:     "Create city error",
		Description: "Failed to create a new city",
		Code:        400,
		CodeStr:     "E_CREATE_FAIL",
	}
	CityDeleteFail = ErrorMessage{
		Message:     "Delete city error",
		Description: "Failed to delete a city",
		Code:        400,
		CodeStr:     "E_DELETE_FAIL",
	}
	CityWrongIdentifier = ErrorMessage{
		Message:     "Wrong city Id",
		Description: "Please provide the city identifier",
		Code:        400,
		CodeStr:     "E_WRONG_IDENTIFIER",
	}
	CityNotFound = ErrorMessage{
		Message:     "City not found",
		Description: "City with provided id was not found",
		Code:        404,
		CodeStr:     "E_CITY_NOT_FOUND",
	}
)
