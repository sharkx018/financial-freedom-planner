package helper

import "math"

func InflationCalculator(amount float64, time int64, rate float64) float64 {
	// Calculate future value including inflation
	futureValue := amount * math.Pow(1+(rate/100), float64(time))
	return RoundToDecimals(futureValue, 2)
}
