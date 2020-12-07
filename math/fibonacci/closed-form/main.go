package main

import (
	"fmt"
	"math"
)

const (
	sqrt5      = 2.236067977499789805051477742381393909454345703125
	plusSqrt5  = (1 + sqrt5) * 0.5
	minusSqrt5 = (1 - sqrt5) * 0.5
	thruSqrt5  = 1 / sqrt5
)

func main() {
	input := float64(0)
	fmt.Print("Enter fibonacci number in sequence to print: ")
	fmt.Scanln(&input)

	pow := math.Pow(plusSqrt5, input) - math.Pow(minusSqrt5, input)
	total := thruSqrt5 * pow
	fmt.Println("The number is:", total)
}
