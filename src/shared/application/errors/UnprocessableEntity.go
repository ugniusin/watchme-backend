package errors

import "net/http"

type UnprocessableEntity struct {
	message string
}

func NewUnprocessableEntity (message string) *UnprocessableEntity {
	return &UnprocessableEntity{
		message: message,
	}
}

func (e *UnprocessableEntity) Error() string {
	return e.message
}

func (e *UnprocessableEntity) GetCode() int {
	return http.StatusUnprocessableEntity
}
