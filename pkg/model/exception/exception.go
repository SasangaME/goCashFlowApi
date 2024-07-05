package exception

type ApplicationError struct {
	IsError      bool
	StatusCode   int
	ErrorMessage string
}

func (e ApplicationError) Error() string {
	return e.ErrorMessage
}
