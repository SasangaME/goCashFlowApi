package dto

type ApplicationResponse struct {
	IsError      bool   `json:"isError"`
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
	StackTrace   string `json:"stackTrace"`
}

func (e ApplicationResponse) Error() string {
	return e.ErrorMessage
}
