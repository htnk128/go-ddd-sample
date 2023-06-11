package resource

type (
	Error struct {
		Type    string `json:"type"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	ErrorResponse struct {
		Error *Error `json:"error"`
	}
)

func NewErrorResponse(errorType string, code int, message string) *ErrorResponse {
	return &ErrorResponse{
		&Error{
			Type:    errorType,
			Status:  code,
			Message: message,
		},
	}
}
