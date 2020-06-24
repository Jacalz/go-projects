package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type mandelbrot struct {
	image         *image.Gray
	width, height int

	complexity float64
	maxIter    int
}

func newMandelbrot(w, h int, c float64) *mandelbrot {
	return &mandelbrot{image: image.NewGray(image.Rect(0, 0, w, h)), width: w, height: h, complexity: c}
}

func (m *mandelbrot) scaleX(a int) float64 {
	return 3.5*float64(a)/float64(m.width) - 1
}

func (m *mandelbrot) scaleY(b int) float64 {
	return 2*float64(b)/float64(m.height) - 1
}

// pixelPainter returns the color of a specified pixel.
func (m *mandelbrot) pixelPainter(a, b int) color.Color {
	x0 := m.scaleX(a)
	y0 := m.scaleY(b)

	var x, y float64

	for i := 0; x*x+y*y <= m.complexity && i < m.maxIter; i++ {
		xtemp := x*x - y*y + x0
		y = 2*x*y + y0
		x = xtemp
	}

	return color.Gray{uint8(x)}
}

func main() {
	mbrot := newMandelbrot(1920, 1080, 4)

	for a := 0; a <= mbrot.width; a++ {
		for b := 0; b <= mbrot.height; b++ {
			mbrot.image.Set(a, b, mbrot.pixelPainter(a, b))
		}
	}

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}

	if err = png.Encode(f, mbrot.image); err != nil {
		panic(err)
	}
}

// for each pixel (Px, Py) on the screen do
// 	x0 = scaled x coordinate of pixel (scaled to lie in the Mandelbrot X scale (-2.5, 1))
// 	y0 = scaled y coordinate of pixel (scaled to lie in the Mandelbrot Y scale (-1, 1))
// 	x := 0.0
// 	y := 0.0
// 	iteration := 0
// 	max_iteration := 1000
// 	while (x×x + y×y ≤ 2×2 AND iteration < max_iteration) do
//     xtemp := x×x - y×y + x0
//     y := 2×x×y + y0
//     x := xtemp
//     iteration := iteration + 1

// 	color := palette[iteration]
// 	plot(Px, Py, color)
