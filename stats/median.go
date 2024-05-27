package stats

import (
	"sort"
)

func Median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)

	// for even number of elements, med is average of two middle values
	if n%2 == 0 {
		med := (data[n/2-1] + data[n/2]) / 2
		return float64(med)
	}
	// for odd number of elements, med is the middle value
	med := data[n/2]
	return float64(med)
}
