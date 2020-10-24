package ambilight

import (
	"fmt"
	"image"
	"log"
	"math"
)

func (s *ambilight) UpdateState() {
	for j, tile := range s.tiles {
		j := j*numberOfColorsInRGB + numberOfHeaderByes
		tile := tile

		pixels, err := s.makeScreenshot(&tile)
		if err != nil {
			log.Println(err)
		}

		r, g, b, err := getAverageColor(pixels)
		if err != nil {
			log.Println(err)
		}

		s.buffer[j+0] = byte(r)
		s.buffer[j+1] = byte(g)
		s.buffer[j+2] = byte(b)

	}

	filterNearest(s.buffer)
}

func filterNearest(buffer []byte) {
	colors := buffer[numberOfHeaderByes:]
	_ = colors
	newBuffer := make([]byte, len(buffer))
	copy(newBuffer, buffer[0:numberOfHeaderByes])

	newBuffer[numberOfHeaderByes+0] = averageColor(colors[len(colors)-numberOfColorsInRGB+0], colors[0], colors[0+numberOfColorsInRGB])
	newBuffer[numberOfHeaderByes+1] = averageColor(colors[len(colors)-numberOfColorsInRGB+1], colors[1], colors[1+numberOfColorsInRGB])
	newBuffer[numberOfHeaderByes+2] = averageColor(colors[len(colors)-numberOfColorsInRGB+2], colors[2], colors[2+numberOfColorsInRGB])

	for
	i := 1* numberOfColorsInRGB;
	i < len(colors)-1*numberOfColorsInRGB;
	i = i + numberOfColorsInRGB {
		r := i + 0
		g := i + 1
		b := i + 2

		newBuffer[r] = averageColor(colors[r-numberOfColorsInRGB], colors[r], colors[r+numberOfColorsInRGB])
		newBuffer[g] = averageColor(colors[g-numberOfColorsInRGB], colors[g], colors[g+numberOfColorsInRGB])
		newBuffer[b] = averageColor(colors[b-numberOfColorsInRGB], colors[b], colors[b+numberOfColorsInRGB])
	}

	newBuffer[len(buffer)-numberOfColorsInRGB+0] = averageColor(colors[len(colors)-2*numberOfColorsInRGB+0], colors[len(colors)-numberOfColorsInRGB+0], colors[0])
	newBuffer[len(buffer)-numberOfColorsInRGB+1] = averageColor(colors[len(colors)-2*numberOfColorsInRGB+1], colors[len(colors)-numberOfColorsInRGB+1], colors[1])
	newBuffer[len(buffer)-numberOfColorsInRGB+2] = averageColor(colors[len(colors)-2*numberOfColorsInRGB+2], colors[len(colors)-numberOfColorsInRGB+2], colors[2])

	buffer = newBuffer
}

func averageColor(previous, current, next byte) (newCurrent byte) {
	newCurrent = byte(math.Round(float64(previous+current+next) / 3))

	return
}

func (s *ambilight) Buffer() []byte {
	return s.buffer
}

func getAverageColor(img *image.RGBA) (r, g, b int, err error) {
	if img == nil {
		fmt.Println("HERE")
	}
	n := len(img.Pix)

	m := 1
	for i := 0; i < n; i = i + 4*m {
		r += int(img.Pix[i])
		g += int(img.Pix[i+1])
		b += int(img.Pix[i+2])
	}

	r = r / n
	g = g / n
	b = b / n

	return
}

func (s *ambilight) makeScreenshot(tl *tile) (img *image.RGBA, err error) {
	img, err = s.screenshoter.CaptureRectangle(image.Rect(tl.x, tl.y, tl.x+tl.width, tl.y+tl.height))
	if err != nil {
		return
	}

	return
}
