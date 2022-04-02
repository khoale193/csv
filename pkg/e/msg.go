package e

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}
	return Msg[ERROR]
}

var Msg = map[int]string{
	SUCCESS_BOOL: "OK",
	ERROR_BOOL:   "FAIL",

	SUCCESS: "success",
	ERR:     "error",

	ERROR_INVALID_NUMBER_OF_COLUMN: "Invalid Number of Column",
	ERROR_CAN_NOT_SAVE_FILE:        "Unable to save the file",
	ERROR_INVALID_UNIX_TIMESTAMP:   "Invalid Unix Timestamp",
}
