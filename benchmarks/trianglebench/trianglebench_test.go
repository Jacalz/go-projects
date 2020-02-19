package trianglebench

import "testing"

func TestRun(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{100, 170},
		{3000, 7608},
		{5000, 13334},
		{50000, 162134},
		{10000, 28394},
	}

	for _, test := range tests {
		if output := Run(test.input); output != test.expected {
			t.Errorf("Test Failed: %v inputed, %v expected, received %v", test.input, test.expected, output)
		}
	}
}
