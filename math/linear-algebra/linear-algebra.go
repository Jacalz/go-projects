package main

import (
	"fmt"
	"math"
)

// Vec3Zero is a zero vector in the room.
var Vec3Zero = &Vec3{0, 0, 0}

// Vec3 is a three dimensional vector in the room.
type Vec3 struct {
	x, y, z float64
}

// Point3 is a point within a three dimensional room.
type Point3 struct {
	x, y, z float64
}

// Vec3BetweenPoints returns a vector from point a to point b.
func Vec3BetweenPoints(a, b *Point3) *Vec3 {
	return &Vec3{b.x - a.x, b.y - a.y, b.z - a.z}
}

// Add adds the two vectors u and v together.
func Add(u, v *Vec3) *Vec3 {
	return &Vec3{u.x + v.x, u.y + v.y, u.z + v.z}
}

// Sub subtracts the vector v from u.
func Sub(u, v *Vec3) *Vec3 {
	return &Vec3{u.x - v.x, u.y - v.y, u.z - v.z}
}

// Mult multiplies the vector u with the scalar s.
func Mult(u *Vec3, s float64) *Vec3 {
	return &Vec3{u.x * s, u.y * s, u.z * s}
}

// Abs returns the absolute value (length) of the vector u.
func Abs(u *Vec3) float64 {
	return math.Sqrt(u.x*u.x + u.y*u.y + u.z*u.z)
}

// ScalarProduct returns the scalar product of the vectors u and v.
func ScalarProduct(u, v *Vec3) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z
}

// CrossProduct returns the cross product of the vectors u and v.
func CrossProduct(u, v *Vec3) *Vec3 {
	return &Vec3{
		x: u.y*v.z - u.z*v.y,
		y: u.z*v.x - v.z*u.x,
		z: u.x*v.y - u.y*v.x,
	}
}

// UnitVector returns a unit vector (length 1) from u.
func UnitVector(u *Vec3) *Vec3 {
	return Mult(u, 1/Abs(u))
}

// OrthoProject projects the vector u orthogonally on the vector v.
func OrthoProject(u, v *Vec3) *Vec3 {
	e := UnitVector(v)
	return Mult(e, ScalarProduct(u, e))
}

// Parallel returns true if the vectors u and v are parallell.
func Parallel(u, v *Vec3) bool {
	return CrossProduct(u, v) == Vec3Zero
}

// Orthogonal returns true if the vectors u and v are orthogonal.
func Orthogonal(u, v *Vec3) bool {
	return ScalarProduct(u, v) == 0
}

// InSamePlane returns true if the vectors u, v and w are in the same plane.
func InSamePlane(u, v, w *Vec3) bool {
	return ScalarProduct(CrossProduct(u, v), w) == 0
}

// AngleBetween returns the angle, in radians, between the vetors u and v.
func AngleBetween(u, v *Vec3) float64 {
	return math.Acos(ScalarProduct(u, v) / (Abs(u) * Abs(v)))
}

func main() {
	u := &Vec3{x: 3, y: 0, z: 1}
	v := &Vec3{x: 2, y: -1, z: 3}
	w := &Vec3{x: 1, y: 1, z: 1}

	fmt.Println("|U|:", Abs(u))

	fmt.Println("UxV:", CrossProduct(u, v))
	fmt.Println("VxU:", CrossProduct(v, u))

	fmt.Println("UxV are parallell:", Parallel(u, v))
	fmt.Println("VxU are parallell:", Parallel(v, u))

	fmt.Println("V, U and W are in the same plane:", InSamePlane(u, v, w))

	p1 := &Point3{x: 1, y: 1, z: 1}
	p2 := &Point3{x: 0, y: 2, z: 1}
	p3 := &Point3{x: -1, y: 0, z: 1}
	p4 := &Point3{x: 2, y: 2, z: -3}

	p1p2 := Vec3BetweenPoints(p1, p2)
	p1p3 := Vec3BetweenPoints(p1, p3)
	p1p4 := Vec3BetweenPoints(p1, p4)

	fmt.Println("p1, p2, p3 and p4 are in the same plane:", InSamePlane(p1p2, p1p3, p1p4))

	u2 := &Vec3{x: 3, y: 0, z: -1}
	v2 := &Vec3{x: 1, y: 2, z: 2}

	fmt.Println("Orthogonal projection of u2 on v2:", OrthoProject(u2, v2))
}
