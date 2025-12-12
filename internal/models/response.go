package models

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Error ErrorResponse `json:"error"`
	Data  interface{}   `json:"data"`
}

func NewSuccessResponse(data interface{}) ApiResponse {
	return ApiResponse{
		Error: ErrorResponse{
			Status:  0,
			Message: "",
		},
		Data: data,
	}
}

func NewErrorResponse(status int, message string) ApiResponse {
	return ApiResponse{
		Error: ErrorResponse{
			Status:  status,
			Message: message,
		},
		Data: nil,
	}
}
