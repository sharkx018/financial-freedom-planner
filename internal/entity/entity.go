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

type AssetClass struct {
	ID                         int64   `json:"id"`                            // bigint corresponds to int64 in Go
	Name                       string  `json:"name"`                          // varchar corresponds to string
	ExpectedReturnInPercentage float64 `json:"expected_return_in_percentage"` // double precision corresponds to float64
}
