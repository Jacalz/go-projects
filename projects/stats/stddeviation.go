package stats

import "math"

// Variance returns the varaiance (average of the varations from the average) of an array of numbers as a float64 number. The second input is a bool and that determens if the data is from a sample. Simples are divided with n - 1 instead of n.
func Variance(numbers []float64, sample bool) (output float64) {
	// Calculate the mean of the array.
	mean := Mean(numbers)

	// Save the length of the array in to a float64 variable.
	length := float64(len(numbers))

	// Create a variable to save the sum.
	var sum float64

	// Loop through and add ((each number in the array minus the average) squared) to the sum.
	for _, num := range numbers {
		sum += math.Pow(num-mean, 2)
	}

	// If we are working with a sample of data, we should divide by length minus one instead.
	if sample {
		output = sum / (length - 1)
	} else {
		output = sum / length
	}

	return output
}

// StdDeviation returns the standard deviation of an array of numbers as float64.
func StdDeviation(numbers []float64, sample bool) float64 {
	return math.Sqrt(Variance(numbers, sample))
}
