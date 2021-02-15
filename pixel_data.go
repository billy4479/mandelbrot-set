package main

import (
	"image/color"
)

type pixelData struct {
	Px, Py     uint64
	x0, y0     float64
	iterations uint64

	color color.Color
}

func newPixel(Px, Py uint64) pixelData {
	return pixelData{
		Px: Px,
		Py: Py,
	}
}

func (d *pixelData) computeIterationCount(maxIterations uint64) {
	// z_{n+1} = z_n^2 + c

	x, y := 0.0, 0.0
	d.iterations = 0

	for x*x+y*y <= 4 && d.iterations < maxIterations {
		xTemp := x*x - y*y + d.x0
		y = 2*x*y + d.y0
		x = xTemp
		d.iterations++
	}

}

func (d *pixelData) computeComplexCoords(xMin, xMax, yMin, yMax float64, width, height uint64) {
	d.x0 = (float64(d.Px) * (xMax - xMin) / float64(width)) + xMin
	d.y0 = (float64(d.Py) * (yMax - yMin) / float64(height)) + yMin
}
