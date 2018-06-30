package domain

type Error struct {
	message string
	err     *Error
}

func (e *Error) Message() string {
	return e.message
}
