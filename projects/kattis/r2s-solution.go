package main

import "fmt"

func main() {
	var R1, S int
	fmt.Scanf("%v\n%v", &R1, &S)
	R2 := 2*S - R1
	fmt.Print(R2)
}
