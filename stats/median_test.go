package stats

import (
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Test1",
			data:     []float64{1, 3, 5},
			expected: 3,
		},
		{
			name:     "Test2",
			data:     []float64{10, 20, 30, 40},
			expected: 25,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Median(test.data)
			if got != test.expected {
				t.Errorf("Median() = %v, want %v", got, test.expected)
			}
		})
	}
}
