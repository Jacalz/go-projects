package main

import "fmt"

// Calculates the values of equation from points on the line
func lineEquation(Point1X, Point1Y, Point2X, Point2Y float64) (Gradient, Intercept float64) {
	Gradient = (Point2Y - Point1Y) / (Point2X - Point1X)
	Intercept = Point1Y - (Gradient * Point1X)
	return Gradient, Intercept
}

func main() {
	var Point1X, Point1Y, Point2X, Point2Y float64
	fmt.Println("Enter two points from straight line graph to get the equation for it!")

	// Enter the first point from the line
	fmt.Print("\nEnter first point (x, y): ")
	fmt.Scanf("(%v, %v)\n", &Point1X, &Point1Y)

	// Enter the second point on the line
	fmt.Print("Enter second point (x, y): ")
	fmt.Scanf("(%v, %v)", &Point2X, &Point2Y)

	// Output values from function to respectiv variables
	Gradient, Intercept := lineEquation(Point1X, Point1Y, Point2X, Point2Y)

	// Get it to not print a + before a negative number
	fmt.Println("\nYou equation is:")
	if Intercept > 0 {
		fmt.Printf("y = %vx + %v", Gradient, Intercept)
	} else if Intercept < 0 {
		Intercept = -Intercept
		fmt.Printf("y = %vx - %v", Gradient, Intercept)
	}

}
