package errors

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	if e.Message == "" {
		return "unknown error"
	}
	return e.Message
}

func (e *CustomError) StatusCode() int {
	return e.Code
}
