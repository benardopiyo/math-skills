package stats

import "math"

func StandardDeviation(data []float64) float64 {
	variance := Variance(data) // Variance() is called to calculate the variance of the data

	stdDev := math.Sqrt(float64(variance))
	return float64(stdDev)
}
