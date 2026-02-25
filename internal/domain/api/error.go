package api

type ErrorDetail struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Error: ErrorDetail{Message: message}}
}
