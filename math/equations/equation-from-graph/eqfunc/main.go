package main

import (
	"flag"
	"fmt"
	"test/quadratic"
	"test/straight"
)

func main() {
	flag.Parse()
	var argument string = flag.Arg(0)

	switch argument {
	case "straight":
		Point1X, Point1Y, Point2X, Point2Y := straight.Input()
		Gradient, Intercept := straight.Equation(Point1X, Point1Y, Point2X, Point2Y)
		straight.Print(Gradient, Intercept)
	case "quadratic":
		Xvalue, Yvalue, NullPoint1, NullPoint2 := quadratic.Input()
		A, B, C := quadratic.Equation(Xvalue, Yvalue, NullPoint1, NullPoint2)
		quadratic.Print(A, B, C)
	default:
		fmt.Println("Usage: eqsolv [equation-type]\nValid equation types are straight and quadratic.")
	}
}
