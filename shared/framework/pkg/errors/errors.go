package errors

func NewAppError(code int, message string, err error) AppError {
	return AppError{Code: code, Message: message, Err: err}
}

func NewAppErrorFromMessage(code int, message string) AppError {
	return AppError{Code: code, Message: message, Err: nil}
}

func NewAppErrorFromCode(code int) AppError {
	return AppError{Code: code, Message: "", Err: nil}
}

type AppError struct {
	Code    int
	Message string
	Err     error
}

// Error implements error.
func (a AppError) Error() string {
	panic("unimplemented")
}
