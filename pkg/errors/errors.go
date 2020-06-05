package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorCode using int to respresent the error
type ErrorCode int

// A list of common expected errors.
const (
	BadRequest       ErrorCode = 400
	ResourceNotFound ErrorCode = 404
	ServerError      ErrorCode = 500
)

// Error specifies the interfaces required by an error in the system.
type Error interface {
	Name() string
	Code() ErrorCode
	Message() string
	Info() map[string]interface{}
	error
	json.Marshaler
}

func NewError(code ErrorCode, message string) Error {
	return &genericError{
		code:    code,
		message: message,
	}
}

func NewBadRequest(err error) Error {
	return &genericError{
		code:    BadRequest,
		message: err.Error(),
	}
}

func NewResourceNotFound(err error) Error {
	return &genericError{
		code:    ResourceNotFound,
		message: err.Error(),
	}
}

// genericError is an implementation of Error that contains
// an code and error message.
type genericError struct {
	code    ErrorCode
	message string
	info    map[string]interface{}
}

func (e *genericError) Name() string {
	return fmt.Sprintf("%v", e.code)
}

func (e *genericError) Code() ErrorCode {
	return e.code
}

func (e *genericError) Message() string {
	return e.message
}

func (e *genericError) Info() map[string]interface{} {
	return e.info
}

func (e *genericError) Error() string {
	return fmt.Sprintf("%v : %v", e.code, e.message)
}

func (e *genericError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string                 `json:"name"`
		Code    ErrorCode              `json:"code"`
		Status  int                    `json:"status"`
		Message string                 `json:"message"`
		Info    map[string]interface{} `json:"info,omitempty"`
	}{e.Name(), e.Code(), e.StatusCode(), e.Message(), e.Info()})
}

func (e *genericError) StatusCode() int {
	httpStatus, ok := map[ErrorCode]int{
		BadRequest:       http.StatusBadRequest,
		ResourceNotFound: http.StatusNotFound,
	}[e.Code()]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}
	return httpStatus
}

func (e *genericError) Fields() map[string]interface{} {
	fields := make(map[string]interface{})
	for k, v := range e.info {
		fields[k] = v
	}
	fields["code"] = e.code
	return fields
}
