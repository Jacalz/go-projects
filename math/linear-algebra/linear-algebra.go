package main

import (
	"fmt"
	"math"
)

// Vec3 is a three dimensional vector
type Vec3 struct {
	x, y, z int
}

// Point3 is a point within a three diemsional space
type Point3 struct {
	x, y, z int
}

// VectorFromPoints returns a vector from point a to point b
func VectorFromPoints(a, b *Point3) *Vec3 {
	return &Vec3{
		x: b.x - a.x,
		y: b.y - a.y,
		z: b.z - a.z,
	}
}

// Abs returns the absolute value of the vector a
func Abs(a *Vec3) float64 {
	return math.Sqrt(float64(a.x*a.x) + float64(a.y*a.y) + float64(a.z*a.z))
}

// Product returns the product of the vectors a and b.
func Product(a, b *Vec3) int {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// CrossProduct returns the cross product of the vectors a and b.
func CrossProduct(a, b *Vec3) *Vec3 {
	return &Vec3{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - b.z*a.x,
		z: a.x*b.y - a.y*b.x,
	}
}

// Parallell states if the vectors a and b are parallell.
func Parallell(a, b *Vec3) bool {
	v := CrossProduct(a, b)
	return v.x == 0 && v.y == 0 && v.z == 0
}

// InSamePlane states if the vectors a, b and c are in the same plane.
func InSamePlane(a, b, c *Vec3) bool {
	return Product(CrossProduct(a, b), c) == 0
}

func main() {
	u := &Vec3{x: 3, y: 0, z: 1}
	v := &Vec3{x: 2, y: -1, z: 3}
	w := &Vec3{x: 1, y: 1, z: 1}

	fmt.Println("|U|:", Abs(u))

	fmt.Println("UxV:", CrossProduct(u, v))
	fmt.Println("VxU:", CrossProduct(v, u))

	fmt.Println("UxV are parallell:", Parallell(u, v))
	fmt.Println("VxU are parallell:", Parallell(v, u))

	fmt.Println("V, U and W are in the same plane:", InSamePlane(u, v, w))

	p1 := &Point3{x: 1, y: 1, z: 1}
	p2 := &Point3{x: 0, y: 2, z: 1}
	p3 := &Point3{x: -1, y: 0, z: 1}
	p4 := &Point3{x: 2, y: 2, z: -3}

	p1p2 := VectorFromPoints(p1, p2)
	p1p3 := VectorFromPoints(p1, p3)
	p1p4 := VectorFromPoints(p1, p4)

	fmt.Println("p1, p2, p3 and p4 are in the same plane:", InSamePlane(p1p2, p1p3, p1p4))
}
