package stats

import "testing"

var (
	first  = []float64{1, 5, 6, 7, 5, 8}
	second = []float64{2, 9, 15, 102, 1}
	third  = []float64{200, 623, 114, 7}
)

func TestTableMean(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected float64
	}{
		{first, 16.0 / 3.0},
		{second, 25.8},
		{third, 236},
	}

	for _, test := range tests {
		if output := Mean(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}

func BenchmarkMean(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Mean(first)
	}
}

func TestTableMedian(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected float64
	}{
		{first, 6.5},
		{second, 15},
		{third, 368.5},
	}

	for _, test := range tests {
		if output := Median(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}

func BenchmarkMedian(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Median(second)
	}
}
