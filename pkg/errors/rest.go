package errors

import "net/http"

type Exception struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int16    `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *Exception) Error() string {
	return e.Message
}

func newException(message, err string, code int16, causes []Causes) *Exception {
	return &Exception{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationException(message string, causes []Causes) *Exception {
	return &Exception{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerErrorException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewUnauthorizedException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func newForbiddenException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewConflictException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Conflict",
		Code:    http.StatusConflict,
	}
}

func NewPreconditionFailedException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Precondition Failed",
		Code:    http.StatusPreconditionFailed,
	}
}

func NewServiceUnavailableException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Service Unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}

func NewTooManyRequestsException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Too Many Requests",
		Code:    http.StatusTooManyRequests,
	}
}

func NewGatewayTimeoutException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Gateway Timeout",
		Code:    http.StatusGatewayTimeout,
	}
}

func NewRequestTimeoutException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Request Timeout",
		Code:    http.StatusRequestTimeout,
	}
}

func NewLengthRequiredException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Length Required",
		Code:    http.StatusLengthRequired,
	}
}

func NewRequestEntityTooLargeException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Request Entity Too Large",
		Code:    http.StatusRequestEntityTooLarge,
	}
}

func NewUnsupportedMediaTypeException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Unsupported Media Type",
		Code:    http.StatusUnsupportedMediaType,
	}
}

func NewUnprocessableEntityException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Unprocessable Entity",
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewFailedDependencyException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Failed Dependency",
		Code:    http.StatusFailedDependency,
	}
}

func NewInsufficientStorageException(message string) *Exception {
	return &Exception{
		Message: message,
		Err:     "Insufficient Storage",
		Code:    http.StatusInsufficientStorage,
	}
}
