package helper

import (
	"encoding/json"
	"math"
	"net/http"
)

func WriteCustomResp(w http.ResponseWriter, headerStatus int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headerStatus)
	json.NewEncoder(w).Encode(response)
	return
}

func RoundToDecimals(value float64, upToPlace int) float64 {
	var x = 1.0
	for upToPlace > 0 {
		x = x * 10
		upToPlace--
	}
	return math.Round(value*x) / x
}
