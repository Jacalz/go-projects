package stats

// Finds if the number exists in the array and returns a bool along with what index that number exists in.
func findNumIndex(value float64, data []float64) (exists bool, index int) {

	// Loop through the whole array.
	for i := 0; i < len(data); i++ {
		// On each iterattion, we check if the seeken value is the same as the value for the given index.
		if value == data[i] {
			// The values are equal and we exit with a true and the index.
			exists, index = true, i
			break
		} else {
			// It wasn't true this time, thus we continue with the next index.
			continue
		}
	}

	return exists, index
}

// LargestIndex returns the index for the largest number in an array of ints.
func largestPosition(numbers []int) (index int) {

	// Start setting largest to the first data value.
	largest := numbers[0]

	// Loop thorough the array, starting at index one.
	for i := 1; i < len(numbers); i++ {

		// Check if largest is smaller than the value of i. If it is, we update largest to that value and set index to that value.
		if largest < numbers[i] {
			largest = numbers[i]
			index = i
		}
	}

	// We lastly return the inrange on arrays and slices provides both the index and value for erange on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.ach entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.dex of the biggest value.
	return index
}

// Mode returns the number that exists the most times in the array of float64 numbers. The first value in the array will be outputed if every number exists once. Currently only work when there is's only one value that appears in most places.
func Mode(numbers []float64) float64 {

	// Create two arrays. One for data with float64 values and one for count using int values.
	data, count := []float64{}, []int{}

	// Create the bool variable for exists and int variable for index.
	var exists bool
	var index int

	// Loop though the input array of numbers.
	for i := 0; i < len(numbers); i++ {
		// Use the finNumIndex function to find index and bool if numbers of i exists in the data slice.
		exists, index = findNumIndex(numbers[i], data)

		// Check if index is zero or if it doesn't exist.
		if i == 0 || !exists {
			// Append the number to the data array.
			data = append(data, numbers[i])

			// Add a one to count on the same index position.
			count = append(count, 1)
		} else {
			// It exists, thus we make the count bigger.
			count[index]++
		}
	}

	// Use largestPosition to find the biggest number in the data array.
	return data[largestPosition(count)]
}
