package stats

import (
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Test1",
			data:     []float64{1, 3, 5},
			expected: 2,
		},
		{
			name:     "Test2",
			data:     []float64{10, 20, 30, 40},
			expected: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := StandardDeviation(test.data)
			if got != test.expected {
				t.Errorf("StandardDeviation() = %v, want %v", got, test.expected)
			}
		})
	}
}
