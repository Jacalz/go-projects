package stats

import "testing"

var (
	forth = []float64{}
	fifth = []float64{}
)

func TestTableMode(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected float64
	}{
		{first, 5},
		{second, 2},
		{third, 200},
	}

	for _, test := range tests {
		if output := Mode(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, recieved %v", test.input, test.expected, output)
		}
	}
}

func BenchmarkMode(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Mode(first)
	}
}
