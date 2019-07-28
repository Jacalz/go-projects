package stats

import (
	"math"
	"testing"
)

func TestTableVariance(t *testing.T) {
	var tests = []struct {
		input    []float64
		boolean  bool
		expected float64
	}{
		{first, false, 4.888888888888889},
		{second, true, 1846.7},
		{third, true, 72796.66666666667},
	}

	for _, test := range tests {
		if output := Variance(test.input, test.boolean); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}

func TestTableStddeviation(t *testing.T) {
	var tests = []struct {
		input    []float64
		boolean  bool
		expected float64
	}{
		{first, false, math.Sqrt(4.888888888888889)},
		{second, true, math.Sqrt(1846.7)},
		{third, true, math.Sqrt(72796.66666666667)},
	}

	for _, test := range tests {
		if output := StdDeviation(test.input, test.boolean); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}
