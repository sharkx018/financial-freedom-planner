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

type Goals struct {
	ID                  int64   `json:"id"`                     // bigint corresponds to int64 in Go
	Name                string  `json:"name"`                   // varchar corresponds to string
	Description         string  `json:"description"`            // double precision corresponds to float64
	YearsLeft           int64   `json:"years_left"`             // double precision corresponds to float64
	InflationPercentage float64 `json:"inflation_percentage"`   // double precision corresponds to float64
	TodayAmount         float64 `json:"today_amount"`           // double precision corresponds to float64
	AllocatedAmount     float64 `json:"allocated_amount"`       // double precision corresponds to float64
	SIPStepUpPercentage float64 `json:"sip_step_up_percentage"` // double precision corresponds to float64
}

type AllocationType struct {
	ID          int64  `json:"id"`          // bigint corresponds to int64 in Go
	Name        string `json:"name"`        // varchar corresponds to string
	Description string `json:"description"` // double precision corresponds to float64
	MinAge      int64  `json:"min_age"`
	MaxAge      int64  `json:"max_age"`
}

type AllocationTypeConfig struct {
	AssetId                int64   `json:"asset_id"`                 // bigint corresponds to int64 in Go
	AssetName              string  `json:"asset_name"`               // bigint corresponds to int64 in Go
	AllocationInPercentage float64 `json:"allocation_in_percentage"` // varchar corresponds to string
}

type InvestableAssetAllocation struct {
	AssetId                int64   `json:"asset_id"`                // bigint corresponds to int64 in Go
	AssetName              string  `json:"asset_name"`              // varchar corresponds to string
	Value                  float64 `json:"value"`                   // varchar corresponds to string
	ContributionPercentage float64 `json:"contribution_percentage"` // varchar corresponds to string
}

type InvestableAssetAllocationAPIResponse struct {
	AssetId   int64             `json:"asset_id"`   // bigint corresponds to int64 in Go
	AssetName string            `json:"asset_name"` // varchar corresponds to string
	Current   ValueContribution `json:"current"`
	Required  ValueContribution `json:"required"`
}

type ValueContribution struct {
	Value                  float64 `json:"value"`                   // varchar corresponds to string
	ContributionPercentage float64 `json:"contribution_percentage"` // varchar corresponds to string
}

type FireResponse struct {
	YearlyExpense           float64
	RetirementYearlyExpense float64
	LeanFire                float64
	Fire                    float64
	FatFire                 float64
	EarlyRetirementAmount   float64
}
