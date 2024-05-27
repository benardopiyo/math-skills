package stats

import "math"

func Variance(data []float64) float64 {
	avg := Average(data)
	sum := 0.0
	n := float64(len(data))

	for _, num := range data {
		sum += math.Pow(num-float64(avg), 2) // calculates the squared diff from the mean, and squares the result
	}
	variance := sum / n
	return float64(variance)
}
