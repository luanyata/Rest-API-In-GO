package errs

import (
	"fmt"
	"net/http"
)

type HttpErrorInterface interface {
	error
	StatusCode() int
	Type() string
}

type httpError struct {
	Message string
	ErrType string
	Status  int
}

var _ HttpErrorInterface = httpError{}

func (e httpError) Error() string {
	return e.Message
}

func (e httpError) StatusCode() int {
	return e.Status
}

func (e httpError) Type() string {
	return e.ErrType
}

func NewHttpError(message, errType string, status int) HttpErrorInterface {
	return httpError{
		Message: message,
		ErrType: errType,
		Status:  status,
	}
}

func BadRequestException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusBadRequest)
}

func UnauthorizedException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusUnauthorized)
}

func NotFoundException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusNotFound)
}

func ForbiddenException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusForbidden)
}

func ConflictException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusConflict)
}

func UnprocessableEntityException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusUnprocessableEntity)
}

func NotAcceptableException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusNotAcceptable)
}

func MethodNotAllowedException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusMethodNotAllowed)
}

func NotImplementedException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusNotImplemented)
}

func ServiceUnavailableException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusServiceUnavailable)
}

func GatewayTimeoutException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusGatewayTimeout)
}

func BadGatewayException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusBadGateway)
}

func TooManyRequestsException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusTooManyRequests)
}

func PreconditionFailedException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusPreconditionFailed)
}

func NotExtendedException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusNotExtended)
}

func ExpectationFailedException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusExpectationFailed)
}

func InternalServerException(message, errType string) HttpErrorInterface {
	return NewHttpError(message, errType, http.StatusInternalServerError)
}

func (e httpError) String() string {
	return fmt.Sprintf("HttpError{status=%d, type=%s, message=%q}", e.Status, e.ErrType, e.Message)
}
