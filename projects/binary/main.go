package main

import (
	"fmt"
	"strconv"
)

func binaryArrayToInt(input []int) int {
	var strnum string
	for i := 0; i < len(input); i++ {
		strnum += strconv.Itoa(input[i])
	}

	output, _ := strconv.Atoi(strnum)
	return output
}

func reverseArray(input []int) (reversed []int) {
	length := len(input)
	for i := 0; i < length; i++ {
		reversed = append(reversed, input[length-(1+i)])
	}

	return reversed
}

func convertToBinary(number int) int {
	var modulus int
	var binary []int

	for number > 0 {
		modulus = number % 2
		number = number / 2
		binary = append(binary, modulus)
	}

	binary = reverseArray(binary)
	return binaryArrayToInt(binary)
}

func printScan(text string) (input int) {
	fmt.Printf("%s: ", text)
	fmt.Scanln(&input)
	return input
}

func main() {
	number := printScan("Please enter an intredger number")
	fmt.Printf("Your binary number is: %v\n", convertToBinary(number))
}
