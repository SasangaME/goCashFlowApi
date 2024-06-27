package exception

type ApplicationError struct {
	StatusCode   int
	ErrorMessage string
}

func (e ApplicationError) Error() string {
	return e.ErrorMessage
}
