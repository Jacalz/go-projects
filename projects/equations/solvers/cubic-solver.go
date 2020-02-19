// This program incorporates the mathematical solving from https://www.1728.org/cubic2.htm

package main

import (
	"fmt"
	"math"
)

func threeRoots(A, B, G, H float64) (X1, X2, X3 float64) {
	// All the stupid complex calculations
	I := math.Sqrt((math.Pow(G, 2) / 4) - H)
	J := math.Cbrt(I)
	K := math.Acos(-(G / (2 * I)))
	L := J * -1
	M := math.Cos(K / 3)
	N := math.Sqrt(3) * math.Sin(K/3)
	P := (B / (3 * A)) * -1

	// Final calculations for X1, X2 and X3
	X1 = (2*J)*math.Cos(K/3) - (B / (3 * A))
	X2 = L*(M+N) + P
	X3 = L*(M-N) + P

	return X1, X2, X3
}

func threeEqualRoots(A, D float64) (X1, X2, X3 float64) {
	X1 = math.Cbrt(D/A) * -1
	X2, X3 = X1, X1
	return X1, X2, X3
}

func oneRealRoot(A, B, C, D, G, H float64) (X1 float64, X2, X3 complex128) {
	R := -(G / 2) + math.Sqrt(H)
	S := math.Cbrt(R)
	T := -(G / 2) - math.Sqrt(H)
	U := math.Cbrt(T)

	X1 = (S + U) - (B / (3 * A))

	Imag := (S - U) * math.Sqrt(3) / 2
	X2 = complex(-(S+U)/2-(B/(3*A)), Imag)
	X3 = complex(-(S+U)/2-(B/(3*A)), -Imag)

	return X1, X2, X3
}

func calculate(A, B, C, D float64) (F, G, H float64) {
	F = (((3 * C) / A) - (math.Pow(B, 2) / math.Pow(A, 2))) / 3
	G = ((2 * math.Pow(B, 3) / math.Pow(A, 3)) - ((9 * B * C) / math.Pow(A, 2)) + ((27 * D) / A)) / 27
	H = (math.Pow(G, 2) / 4) + (math.Pow(F, 3) / 27)

	return F, G, H
}

func getInput() (A, B, C, D float64) {
	fmt.Println("Enter values for equation in format ax³ + bx² + cx + d = 0:")
	fmt.Print("Value of A: ")
	fmt.Scanln(&A)
	fmt.Print("value of B: ")
	fmt.Scanln(&B)
	fmt.Print("value of C: ")
	fmt.Scanln(&C)
	fmt.Print("value of D: ")
	fmt.Scanln(&D)

	return A, B, C, D
}

func main() {
	A, B, C, D := getInput()
	F, G, H := calculate(A, B, C, D)

	var X1, X2, X3 float64
	var X2I, X3I complex128
	if H > 0 {
		X1, X2I, X3I = oneRealRoot(A, B, C, D, G, H)
	} else if F == 0 && G == 0 && H == 0 {
		X1, X2, X3 = threeEqualRoots(A, D)
	} else if H <= 0 {
		X1, X2, X3 = threeRoots(A, B, G, H)
	}

	if H > 0 {
		fmt.Println("\nSolutions for equation:")
		fmt.Printf("X1: %.3f\n", X1)
		fmt.Printf("X2: %.3f\n", X2I)
		fmt.Printf("X3: %.3f\n", X3I)
	} else {
		fmt.Println("\nSolutions for equation:")
		fmt.Printf("X1: %.3f\n", X1)
		fmt.Printf("X2: %.3f\n", X2)
		fmt.Printf("X3: %.3f\n", X3)
	}

}
