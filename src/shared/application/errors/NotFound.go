package errors

import "net/http"

type NotFound struct {
	message string
}

func NewNotFound (message string) *NotFound {
	return &NotFound{
		message: message,
	}
}

func (e *NotFound) Error() string {
	return e.message
}

func (e *NotFound) GetCode() int {
	return http.StatusNotFound
}
