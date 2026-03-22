package errors

type ErrorResponse struct {
	Error ErrorResponseError `json:"error"`
}