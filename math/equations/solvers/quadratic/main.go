package main

import (
	"fmt"
	"log"
	"math"
)

func complexRoots(B, root float64) (X1I, X2I complex128) {
	X1I = complex(-(B / 2), +math.Sqrt(-root))
	X2I = complex(-(B / 2), -math.Sqrt(-root))

	return X1I, X2I
}

func realRoots(B, root float64) (X1, X2 float64) {
	X1 = -(B / 2) + math.Sqrt(root)
	X2 = -(B / 2) - math.Sqrt(root)

	return X1, X2
}

func rootPart(A, B, C float64) (root float64) {
	root = math.Pow(B/2, 2) - C

	return root
}

func getInput() (A, B, C float64) {
	fmt.Println("Enter values for equation in format ax² + bx + x = 0:")
	fmt.Print("Value of A: ")
	fmt.Scanln(&A)
	fmt.Print("value of B: ")
	fmt.Scanln(&B)
	fmt.Print("value of C: ")
	fmt.Scanln(&C)

	return A, B, C
}

func handleInputValues(A, B, C float64) (AO, BO, CO float64) {
	if A != 1 {
		AO = 1
		BO = B / A
		CO = C / A
	} else if A == 1 {
		AO = A
		BO = B
		CO = C
	} else if A == 0 {
		log.Fatalln("Not a quadratic equation")
	}

	return AO, BO, CO
}

func main() {

	A, B, C := handleInputValues(getInput())
	root := rootPart(A, B, C)

	fmt.Println("\nSolutions for equation:")

	fmt.Println("Debug:", B, root)

	if root < 0 {
		X1I, X2I := complexRoots(B, root)
		fmt.Printf("X1: %.3f\n", X1I)
		fmt.Printf("X2: %.3f\n", X2I)
	} else {
		X1, X2 := realRoots(B, root)
		fmt.Printf("X1: %.3f\n", X1)
		fmt.Printf("X2: %.3f\n", X2)
	}

}
