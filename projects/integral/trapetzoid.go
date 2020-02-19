package main

import "fmt"

// Define the equation to integrate.
func f(x float64) float64 {
	return x * x
}

// Integral gives us the area under a graph between the start and end values. The accuracy is defined by telling the amount of trapzoids to divide it up into.
func integral(start, end float64, amount int) float64 {

	// h gives us the length of each divided part.
	h := (end - start) / float64(amount)

	// s is our intital sum of half the start values.
	s := 0.5 * (f(start) + f(end))

	// Loop through the sums of the functions at each trapzoid.
	for i := 1; i < amount-1; i++ {
		s += f(start + float64(i)*h)
	}

	// The area is our h multiplied by the sum.
	I := h * s

	return I
}

func main() {
	fmt.Println(integral(2, 4, 1000000))
}
