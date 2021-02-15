package main

import (
	"image/color"
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
)

const (
	BLACK_AND_WHITE = iota
	HUE
)

func (d *pixelData) computeColor(maxIterations uint64, colorMode int) {
	switch colorMode {
	case BLACK_AND_WHITE:
		var c uint8
		if d.iterations == maxIterations {
			c = 0
		} else {
			c = uint8(
				math.Round(
					float64(maxIterations) * math.Sqrt(
						float64(d.iterations)/float64(maxIterations),
					),
				),
			)
		}
		d.color = color.RGBA{
			R: c, G: c, B: c, A: 0xff,
		}
		break
	case HUE:
		hue := 360 * float64(d.iterations) / float64(maxIterations)
		saturation := 1
		value := 1
		if d.iterations == maxIterations {
			value = 0
		}

		d.color = colorful.Hsv(hue, float64(saturation), float64(value))
	}
}
