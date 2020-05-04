package swirl

import "math"

type swirl struct {
	buffer []byte
	sine1  float64
	sine2  float64
	hue1   int
	hue2   int
}

const (
	numberOfColorsInRGB = 3
	numberOfHeaderByes  = 6
)

func New(nLeds int) *swirl {

	size := numberOfHeaderByes + nLeds*numberOfColorsInRGB

	buffer := make([]byte, size)
	buffer[0] = 'A' // Magic word
	buffer[1] = 'd'
	buffer[2] = 'a'
	buffer[3] = byte((nLeds - 1) >> 8)             // LED count high byte
	buffer[4] = byte((nLeds - 1) & 0xff)           // LED count low byte
	buffer[5] = byte(buffer[3] ^ buffer[4] ^ 0x55) // Checksum

	return &swirl{
		buffer: buffer,
		sine1:  0,
		sine2:  0,
		hue1:   0,
		hue2:   0,
	}
}

func (s *swirl) Buffer() []byte {
	return s.buffer
}

func (s *swirl) UpdateState() {

	for i := 6; i < len(s.buffer); {
		// Fixed-point hue-to-RGB conversion.  'hue2' is an integer in the
		// range of 0 to 1535, where 0 = red, 256 = yellow, 512 = green, etc.
		// The high byte (0-5) corresponds to the sextant within the color
		// wheel, while the low byte (0-255) is the fractional part between
		// the primary/secondary colors.
		lo := s.hue2 & 255
		var r, g, b int
		switch (s.hue2 >> 8) % 6 {
		case 0:
			r = 255
			g = lo
			b = 0
		case 1:
			r = 255 - lo
			g = 255
			b = 0
		case 2:
			r = 0
			g = 255
			b = lo
		case 3:
			r = 0
			g = 255 - lo
			b = 255
		case 4:
			r = lo
			g = 0
			b = 255
		default:
			r = 255
			g = 0
			b = 255 - lo
		}

		// Resulting hue is multiplied by brightness in the range of 0 to 255
		// (0 = off, 255 = brightest).  Gamma correction (the 'pow' function
		// here) adjusts the brightness to be more perceptually linear.
		c1 := 0.5
		c2 := 0.5
		c3 := 1.8
		bright := int(math.Pow(0.5+math.Sin(s.sine2)*0.5, 1.8) * 255.0)
		s.buffer[i] = byte((r * bright) / 255)
		i++
		s.buffer[i] = byte((g * bright) / 255)
		i++
		s.buffer[i] = byte((b * bright) / 255)
		i++
		// Each pixel is slightly offset in both hue and brightness
		s.hue2 += 40
		s.sine2 += 0.3
	}

	s.hue1 = (s.hue1 + 4) % 1536
	s.sine1 -= .13

}
