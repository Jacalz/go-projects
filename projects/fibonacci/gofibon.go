package main

import "fmt"

func main() {
	var num1, num2, nextTerm int64
	var n int
	num1 = 0
	num2 = 1
	fmt.Print("Enter number of passes: ")
	fmt.Scanln(&n)
	fmt.Println("Fibonacci series:")

	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Println(num1)
			continue
		}
		if i == 2 {
			fmt.Println(num2)
			continue
		}
		nextTerm = num1 + num2
		num1 = num2
		num2 = nextTerm
		fmt.Println(nextTerm)
	}
}
