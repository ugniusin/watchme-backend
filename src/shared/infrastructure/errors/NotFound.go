package errors

type NotFound struct {
	message string
}

func NewError (message string) *NotFound {
	return &NotFound{
		message: message,
	}
}

func (e *NotFound) Error() string { return e.message }
