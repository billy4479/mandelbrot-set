package main

import (
	"image"
	"image/png"
	"os"
	"sync"
)

const (
	width  uint64 = 1920
	height uint64 = 1080
	// width  uint64  = 3840
	// height uint64  = 2160
	r    float64 = 0.002
	xMin float64 = -0.374004139 - r/2 + 0.0005
	xMax float64 = xMin + r/2
	yMin float64 = -0.659792175 - r/2 + 0.0005
	yMax float64 = yMin + r/2
	// xMin          float64 = -0.7463
	// xMax          float64 = xMin + 0.005
	// yMin          float64 = 0.1102
	// yMax          float64 = yMin + 0.005
	maxIterations uint64 = 1000
)

func main() {
	img := mandelbrotSet(xMin, xMax, yMin, yMax, width, height, maxIterations)

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

type img struct {
	img *image.RGBA
	sync.Mutex
}

func mandelbrotSet(xMin, xMax, yMin, yMax float64, width, height, maxIterations uint64) image.Image {
	result := img{
		img: image.NewRGBA(
			image.Rect(0, 0, int(width), int(height)),
		),
	}
	var wg sync.WaitGroup

	// For each pixel on the screen
	var Px, Py uint64
	for Px = 0; Px < width; Px++ {
		for Py = 0; Py < height; Py++ {
			wg.Add(1)
			go func(xMin, xMax, yMin, yMax float64, width, height, maxIterations, x, y uint64) {
				p := newPixel(x, y)
				p.computeComplexCoords(xMin, xMax, yMin, yMax, width, height)
				p.computeIterationCount(maxIterations)
				p.computeColor(maxIterations, HUE)

				result.Lock()
				result.img.Set(int(p.Px), int(p.Py), p.color)
				result.Unlock()
				wg.Done()
			}(xMin, xMax, yMin, yMax, width, height, maxIterations, Px, Py)
		}
	}

	wg.Wait()
	return result.img
}
