package ambilight

import (
	"fmt"
	"image"
	"log"
	"sync"
)

func (s *ambilight) UpdateState() {
	var wg sync.WaitGroup
	wg.Add(len(s.tiles))

	for j, tile := range s.tiles {
		j := j*numberOfColorsInRGB + numberOfHeaderByes
		tile := tile

		go func() {
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
			wg.Done()
		}()
	}

	wg.Wait()
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
