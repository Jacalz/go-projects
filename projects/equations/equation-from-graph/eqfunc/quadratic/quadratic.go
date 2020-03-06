package quadratic

import (
	"fmt"
	"log"
)

// Input will get all the values needed to calculate the quadratic equation
func Input() (Xvalue, Yvalue, NullPoint1, NullPoint2 float64) {
	// Enter the amount of roots in the graph
	var NullPoints int

	fmt.Print("\nAmount of roots on the graph (points where y = 0): ")
	fmt.Scanln(&NullPoints)
	NullPoint1, NullPoint2 = equationRoots(NullPoints)

	// A random value from the graph to calculate gradient of line
	fmt.Print("\nEnter a given point on the graph (x, y): ")
	fmt.Scanf("(%v, %v)", &Xvalue, &Yvalue)

	return Xvalue, Yvalue, NullPoint1, NullPoint2
}

// Handle different amounts of roots for the graph
func equationRoots(NullPoints int) (NullPoint1, NullPoint2 float64) {
	switch NullPoints {
	case 2:
		NullPoint1, NullPoint2 = twoRoots()
	case 1:
		NullPoint1 = oneRoot()
		// The point is called a rebound point and two root points will have the same value
		NullPoint2 = NullPoint1
	case 0:
		log.Fatalln("The function has no real roots and can't be parsed by this program!")
	default:
		log.Fatalln("A quadratic function can only have 2, 1 or 0 roots!")
	}
	return
}

// Get root points from graph when there are two possible roots
func twoRoots() (NullPoint1, NullPoint2 float64) {
	fmt.Println("\nEnter x-values for the two root points:")
	fmt.Print("First x-value: ")
	fmt.Scan(&NullPoint1)

	fmt.Print("Second x-value: ")
	fmt.Scan(&NullPoint2)

	return NullPoint1, NullPoint2
}

// Get the root point for one root
func oneRoot() (NullPoint float64) {
	fmt.Println("\nEnter x-value for the root point:")
	fmt.Print("X-value: ")
	fmt.Scan(&NullPoint)

	return NullPoint
}

// Equation calculates the A, B and C value for the equation
func Equation(Xvalue, Yvalue, NullPoint1, NullPoint2 float64) (A, B, C float64) {
	A = Yvalue / ((Xvalue - NullPoint1) * (Xvalue - NullPoint2))
	B = A * (-NullPoint1 + -NullPoint2)
	C = A * (NullPoint1 * NullPoint2)
	return A, B, C
}

// Print prints the full equation for the graph, from A, B and C values from QuadraticEquation
func Print(A, B, C float64) {
	// Make sure that it doesn't print a plus sign before a negative number
	fmt.Println("\nYou equation is:")
	if B > 0 && C > 0 {
		fmt.Printf("y = %vx² + %vx + %v\n", A, B, C)
	} else if B < 0 && C < 0 {
		B, C = -B, -C
		fmt.Printf("y = %vx² - %vx - %v\n", A, B, C)
	} else if B > 0 && C < 0 {
		C = -C
		fmt.Printf("y = %vx² + %vx - %v\n", A, B, C)
	} else if B < 0 && C > 0 {
		B = -B
		fmt.Printf("y = %vx² - %vx + %v\n", A, B, C)
	}
}
