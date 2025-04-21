package helper

import "math"

func InflationCalculator(amount float64, time int64, rate float64) float64 {
	// Calculate future value including inflation
	futureValue := amount * math.Pow(1+(rate/100), float64(time))
	return RoundToDecimals(futureValue, 2)
}
func CalculateSIPRequired(targetAmount float64, years int64, annualGrowthRate float64, stepUpPercentage float64) float64 {
	monthlyRate := annualGrowthRate / (12 * 100)
	totalMonths := years * 12

	denominator := 0.0

	for i := 0; i < int(totalMonths); i++ {
		currentYear := i / 12
		stepUpMultiplier := math.Pow(1+stepUpPercentage/100, float64(currentYear))
		monthsLeft := int(totalMonths) - i

		installmentFV := stepUpMultiplier * math.Pow(1+monthlyRate, float64(monthsLeft))
		denominator += installmentFV
	}

	sipAmount := targetAmount / denominator

	return RoundToDecimals(sipAmount, 2)
}
