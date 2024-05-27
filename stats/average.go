package stats

func Average(data []float64) float64 {
	sum := 0.0
	n := float64(len(data))

	for _, num := range data {
		sum += num
	}
	avg := sum / n
	return float64(avg)
}
