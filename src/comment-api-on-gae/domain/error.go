package domain

type Error struct {
	message string
	error   *Error
}

func (e *Error) Error() string {
	if e.message != "" {
		return e.message
	} else if e.error != nil {
		return e.Error()
	} else {
		return ""
	}
}
