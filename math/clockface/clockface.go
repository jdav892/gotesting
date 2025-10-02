package clockface

import (
	"time"
)

// A Point represents a 2d cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the seonc hand of the analogue clock at time 't'
// represented as a point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
