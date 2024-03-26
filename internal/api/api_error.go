package handlers

type APIError struct {
	err        error
	statusCode int
}

func (e *APIError) Error() string {
	return e.Error()
}
