package measurer

import (
	"fmt"
	"time"
)

type measurer struct {
	currentTime time.Time
	timePassed  time.Duration
	counter     int64
}

func New() *measurer {

	return &measurer{
		currentTime: time.Now(),
		timePassed:  0,
		counter:     0,
	}
}

func (m *measurer) Print() {
	m.counter++
	m.timePassed = time.Since(m.currentTime)

	if m.timePassed > time.Second {
		tIteration := float32(m.timePassed.Milliseconds()) / float32(m.counter)
		//fmt.Printf("\033[0;100H")
		fmt.Printf("FPS: %.1f\tt: %.1f мс\n", 1000./tIteration, tIteration)

		m.currentTime = time.Now()
		m.counter = 0
		m.timePassed = 0
	}
}
