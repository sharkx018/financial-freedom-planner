package entity

var jwtKey = []byte("my_secret_key")

type ApiResponse struct {
	Data    interface{}          `json:"data"`
	Success bool                 `json:"success"`
	Error   *CommonErrorResponse `json:"error,omitempty"`
}

type CommonErrorResponse struct {
	Message string `json:"message"`
}
