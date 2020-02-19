package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string = "5x^3 - 5x^2 + 9x + 6"

	// Slit up the input in to array for evrything between spaces.
	splited := strings.Split(input, " ")

	// Declare variabes needed in our loop.
	var num1, num2 int
	var output string

	// Range through all the objects in our array.
	for i := range splited {
		// Check if it contains "x^", a number or just "x".
		if strings.Contains(splited[i], "x^") {

			// Take out numbers from the object stored in the array.
			fmt.Sscanf(splited[i], "%vx^%v", &num1, &num2)

			// Calculate the value changes and add them back to the string.
			splited[i] = fmt.Sprintf("%vx^%v", num1*num2, num2-1)

			// Replace "^1" with "" since "x^1" == "x".
			splited[i] = strings.ReplaceAll(splited[i], "^1", "")

			// Check if the objet is a number.
		} else if _, err := strconv.Atoi(splited[i]); err == nil {
			// Remove that value and remove the symbol in front.
			splited[i], splited[i-1] = "", ""
		} else if strings.Contains(splited[i], "x") {

			// Replace all x values with ones if we only have x, otherwice remove the x.
			if splited[i] == "x" {
				splited[i] = strings.ReplaceAll(splited[i], "x", "1")
			} else {
				splited[i] = strings.ReplaceAll(splited[i], "x", "")
			}
		}
	}

	for i := range splited {
		output += splited[i] + " "
	}

	fmt.Println("Din deriverade funktion är:", output)
}
