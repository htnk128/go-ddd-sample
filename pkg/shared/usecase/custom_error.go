package usecase

import (
	"fmt"

	"github.com/friendsofgo/errors"
)

type customError struct {
	errType string
	status  int
	message string
}

const (
	InvalidRequestErrorType   = "invalid_request_error"
	NotFoundErrorType         = "not_found_error"
	InvalidDataStateErrorType = "invalid_data_state"
	UpdateFailureErrorType    = "update_failure_error"
	ServerErrorType           = "server_error"
)

const (
	InvalidRequestErrorStatus   = 400
	NotFoundErrorStatus         = 404
	InvalidDataStateErrorStatus = 409
	UpdateFailureErrorStatus    = 500
	ServerErrorStatus           = 500
)

func NewInvalidRequestError(message string) error {
	return customError{
		errType: InvalidRequestErrorType,
		status:  InvalidRequestErrorStatus,
		message: message,
	}
}

func NewNotFoundError(message string) error {
	return customError{
		errType: NotFoundErrorType,
		status:  NotFoundErrorStatus,
		message: message,
	}
}

func NewInvalidDataStateError(message string) error {
	return customError{
		errType: InvalidDataStateErrorType,
		status:  InvalidDataStateErrorStatus,
		message: message,
	}
}

func NewUpdateFailureError(message string) error {
	return customError{
		errType: UpdateFailureErrorType,
		status:  UpdateFailureErrorStatus,
		message: message,
	}
}

func NewServerError(message string) error {
	return customError{
		errType: ServerErrorType,
		status:  ServerErrorStatus,
		message: message,
	}
}

func (e customError) Error() string {
	return fmt.Sprintf("errType: %s, status: %v, Message: %s", e.errType, e.status, e.message)
}

func ErrorType(err error) string {
	if err == nil {
		return ""
	}

	ok, e := customErrorFromError(err)
	if ok {
		return e.errType
	}

	return ServerErrorType
}

func ErrorStatus(err error) int {
	if err == nil {
		return 0
	}

	ok, e := customErrorFromError(err)
	if ok {
		return e.status
	}

	return ServerErrorStatus
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	ok, e := customErrorFromError(err)
	if ok {
		return e.message
	}

	return "Internal Server Error"
}

func customErrorFromError(err error) (bool, customError) {
	var e customError

	ok := errors.As(err, &e)
	if ok {
		return true, e
	}
	return false, e
}
