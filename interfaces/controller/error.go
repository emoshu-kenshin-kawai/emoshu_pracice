package controller

type Error struct {
	Status  uint
	Message string
}

func NewError(status uint, err error) *Error {
	return &Error{
		Status:  status,
		Message: err.Error(),
	}
}
