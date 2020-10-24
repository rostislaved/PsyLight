package measurer

import (
	"fmt"
	"time"
)

type measurer struct {
	currentTime   time.Time
	timePassed    time.Duration
	counter       int64
	meanFPS       float64
	numOfMeasures float64
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
		timeOfCurrentIteration := float64(m.timePassed.Milliseconds()) / float64(m.counter)

		currentFPS := 1000. / timeOfCurrentIteration
		m.meanFPS = getMeanFPS(m.meanFPS, m.numOfMeasures, currentFPS)
		m.numOfMeasures++

		fmt.Printf("\rFPS: %.1f\tt: %.1f ms \tmeanFPS: %.1f", currentFPS, timeOfCurrentIteration, m.meanFPS)

		m.counter = 0
		m.timePassed = 0
		m.currentTime = time.Now()
	}
}

func getMeanFPS(meanFPS, numOfMeasures, newFPS float64) float64 {
	return (meanFPS*numOfMeasures + newFPS) / (numOfMeasures + 1)
}
