package dto

// APIResponse represents standard successful API response structure
type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data,omitempty"`
}

// APIErrorResponse represents standard error API response structure
type APIErrorResponse struct {
	Success bool     `json:"success" example:"false"`
	Message string   `json:"message" example:"An error occurred"`
	Errors  []string `json:"errors,omitempty" example:"[\"Invalid input\"]"`
}

// NewSuccessResponse creates a new success API response
func NewSuccessResponse(message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// NewErrorResponse creates a new error API response
func NewErrorResponse(message string, errors ...string) APIErrorResponse {
	var errList []string
	if len(errors) > 0 {
		errList = errors
	}
	return APIErrorResponse{
		Success: false,
		Message: message,
		Errors:  errList,
	}
}
