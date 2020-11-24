package main

import "fmt"

func main() {
	var n1, n2 uint = 0, 1
	var next, iter uint

	fmt.Print("Enter amount of numbers to print: ")
	fmt.Scanln(&iter)

	for i := uint(0); i < iter; i++ {
		fmt.Println(next)
		next = n1 + n2
		n2 = n1
		n1 = next
	}
}

