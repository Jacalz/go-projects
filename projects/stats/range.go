package stats

// Largest returns the biggest value from the data in the array of float64 numbers.
func Largest(numbers []float64) float64 {
	// Start setting largest to the first data value.
	largest := numbers[0]

	// Loop though the numbers on an index of one.
	for i := 1; i < len(numbers); i++ {
		// If largest is smaller than numbers for i, we update largest to that value.
		if largest < numbers[i] {
			largest = numbers[i]
		}
	}

	// Lastly, we return the largest value.
	return largest
}

// Smallest returns the smallest value from the data in the array of float64 numbers.
func Smallest(numbers []float64) float64 {
	// Start setting smallest to the first data value.
	smallest := numbers[0]

	// Loop through all our numbers from an index of one.
	for i := 1; i < len(numbers); i++ {
		// If smallest is bigger than numbers for i, we update largest to that value.
		if smallest > numbers[i] {
			smallest = numbers[i]
		}
	}

	// Lastly, we return the smallest value.
	return smallest
}

// Range retuns the length between the biggest and smallest values from an array of float64 numbers.
func Range(numbers []float64) float64 {
	return Largest(numbers) - Smallest(numbers)
}
