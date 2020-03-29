package straight

import (
	"fmt"
	"log"
)

// Input gets all the needed values needed to calculate a straight line equation
func Input() (Point1X, Point1Y, Point2X, Point2Y float64) {
	fmt.Println("Enter two points from straight line graph to get the equation for it!")

	// Enter the first point from the line
	fmt.Print("\nEnter first point (x, y): ")
	fmt.Scanf("(%v, %v)\n", &Point1X, &Point1Y)

	// Enter the second point on the line
	fmt.Print("Enter second point (x, y): ")
	fmt.Scanf("(%v, %v)", &Point2X, &Point2Y)

	if Point1X == Point2X || Point1Y == Point2Y {
		log.Fatalln("The specified points need to be different!")
	}

	return Point1X, Point1Y, Point2X, Point2Y
}

// Equation calculates the gradient and intercept points for the equation
func Equation(Point1X, Point1Y, Point2X, Point2Y float64) (Gradient, Intercept float64) {
	Gradient = (Point2Y - Point1Y) / (Point2X - Point1X)
	Intercept = Point1Y - (Gradient * Point1X)
	return Gradient, Intercept
}

// Print prints the equation for the graph, from gradient and intercept value from StraightEquation
func Print(Gradient, Intercept float64) {
	// Get it to not print a + before a negative number
	fmt.Println("\nYou equation is:")
	if Intercept > 0 {
		fmt.Printf("y = %vx + %v\n", Gradient, Intercept)
	} else if Intercept < 0 {
		Intercept = -Intercept
		fmt.Printf("y = %vx - %v\n", Gradient, Intercept)
	}
}
