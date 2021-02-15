package main

import (
	"flag"
	"image"
	"image/png"
	"os"
	"sync"
)

var (
	width         *uint64  = flag.Uint64("width", 1920, "Image width")
	height        *uint64  = flag.Uint64("height", 1080, "Image height")
	xMin          *float64 = flag.Float64("xMin", -2.5, "Minimum value of X painted on the image")
	xMax          *float64 = flag.Float64("xMaz", 1, "Maximum value of X painted on the image")
	yMin          *float64 = flag.Float64("yMin", -1, "Minimum value of X painted on the image")
	yMax          *float64 = flag.Float64("yMax", 1, "Maximum value of X painted on the image")
	maxIterations *uint64  = flag.Uint64("maxIterations", 100, "Maximum iteration count")
	col           *int     = flag.Int("color", 1, "Color mode (use --help to see the available modes)")

	out  *string = flag.String("out", "output.png", "The output image, maut be a png file")
	help *bool   = flag.Bool("help", false, "Print a help message and exit")
)

func main() {
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	img := mandelbrotSet(*xMin, *xMax, *yMin, *yMax, *width, *height, *maxIterations, *col)

	file, err := os.Create(*out)
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

func mandelbrotSet(xMin, xMax, yMin, yMax float64, width, height, maxIterations uint64, col int) image.Image {
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
			go func(xMin, xMax, yMin, yMax float64, width, height, maxIterations, x, y uint64, col int) {
				p := newPixel(x, y)
				p.computeComplexCoords(xMin, xMax, yMin, yMax, width, height)
				p.computeIterationCount(maxIterations)
				p.computeColor(maxIterations, col)

				result.Lock()
				result.img.Set(int(p.Px), int(p.Py), p.color)
				result.Unlock()
				wg.Done()
			}(xMin, xMax, yMin, yMax, width, height, maxIterations, Px, Py, col)
		}
	}

	wg.Wait()
	return result.img
}
