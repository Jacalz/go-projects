package main

import "fmt"

type Vec3 struct {
	x, y, z int
}

func CrossProduct(a, b *Vec3) *Vec3 {
	return &Vec3{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - b.z*a.x,
		z: a.x*b.y - a.y*b.x,
	}
}

func main() {
	u := &Vec3{x: 4, y: 2, z: 1}
	v := &Vec3{x: 6, y: -8, z: 12}

	fmt.Println("UxV:", CrossProduct(u, v))
	fmt.Println("VxU:", CrossProduct(v, u))
}
