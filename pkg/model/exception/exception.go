package exception

type ApplicationStatus struct {
	StatusCode   int
	ErrorMessage string
}

func (e ApplicationStatus) Error() string {
	return e.ErrorMessage
}
