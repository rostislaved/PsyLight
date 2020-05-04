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

	if m.timePassed >= time.Second {
		tIteration := float64(m.timePassed.Milliseconds()) / float64(m.counter)
		fmt.Printf("FPS: %.1f\tt: %.1f ms\n", 1000./tIteration, tIteration)

		m.counter = 0
		m.timePassed = 0
		m.currentTime = time.Now()
	}
}
