package main

import "fmt"

func gradient(YaxisX, YaxisY, NullPoint1, NullPoint2 float64) float64 {
	return YaxisY / ((YaxisX - NullPoint1) * (YaxisX - NullPoint2))
}

func main() {
	var YaxisX, YaxisY, NullPoint1, NullPoint2 float64
	fmt.Println("Enter a couple values from graph to get quadratic equation!")

	fmt.Print("\nEnter point where line goes through y-axis (x, y): ")
	fmt.Scanf("(%v, %v)", &YaxisX, &YaxisY)

	fmt.Println("\nEnter x-values for points where y-value is zero:")
	fmt.Print("First x-value: ")
	fmt.Scan(&NullPoint1)

	fmt.Print("Second x-value: ")
	fmt.Scan(&NullPoint2)

	Gradient := gradient(YaxisX, YaxisY, NullPoint1, NullPoint2)
	B := Gradient * (-NullPoint1 + -NullPoint2)
	C := Gradient * (NullPoint1 * NullPoint2)

	fmt.Println("\nYou equation is:")
	if B > 0 && C > 0 {
		fmt.Printf("%vx² + %vx + %v = 0\n", Gradient, B, C)
	} else if B < 0 && C < 0 {
		B, C = -B, -C
		fmt.Printf("%vx² - %vx - %v = 0\n", Gradient, B, C)
	} else if B > 0 && C < 0 {
		C = -C
		fmt.Printf("%vx² + %vx - %v = 0\n", Gradient, B, C)
	} else if B < 0 && C > 0 {
		B = -B
		fmt.Printf("%vx² - %vx + %v = 0\n", Gradient, B, C)
	}

}