package stats

import "math"

// Variance returns the varaiance (average of the varations from the average) of an array of numbers as a float64 number. The second input is a bool and that determens if the data is from a sample. Simples are divided with n - 1 instead of n.
func Variance(numbers []float64, sample bool) float64 {
	// Calculate the Mean of the array.
	average := Mean(numbers)

	// Save the length of the array in to a variable.
	length := len(numbers)

	// Create a variable to save the sum.
	var sum float64

	// Loop through and add ((each number in the array minus the average) squared) to the sum.
	for i := 0; i < length; i++ {
		sum += math.Pow(numbers[i]-average, 2)
	}

	// Assign an output variable.
	var output float64

	// If we are owkring with a sample of data, we act accordingly.
	if sample {
		output = sum / (float64(length) - 1)
	} else {
		output = sum / float64(length)
	}

	// We lastly return the output.
	return output
}

// StdDeviation returns the standard deviation of an array of numbers as float64.
func StdDeviation(numbers []float64, sample bool) float64 {
	return math.Sqrt(Variance(numbers, sample))
}
