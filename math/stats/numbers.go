package stats

// Mean takes an array of float64 and returns the average as a float64 number.
func Mean(numbers []float64) float64 {
	// Assign a variable to hold the sum.
	var sum float64

	// Loop through each value and add it to the sum.
	for _, num := range numbers {
		sum += num
	}

	// Return the sum divided by the float64 value of the amount of numbers.
	return sum / float64(len(numbers))
}

// Median returns the number in the middle of the array of float64 numbers.
func Median(numbers []float64) (output float64) {
	// Assign the length of the arrays to get the number of values.
	length := len(numbers)

	// Add an index variable for specifying the index of the value.
	var index float64

	// Check if the array has an equal number of values or not. handle accordingly.
	if length%2 == 0 {
		// The index in this case is will be half of the length.
		index = float64(length) * 0.5

		// The output should be the mean of the middle values.
		output = (numbers[int(index)] + numbers[int(index-1)]) * 0.5
	} else {
		// The index is half the length plus 0.5 to get an even number.
		index = 0.5*float64(length) + 0.5

		// Output should be the number with the index.
		output = numbers[int(index-1)]
	}

	return output
}
