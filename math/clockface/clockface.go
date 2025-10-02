package clockface

import (
	"time"
)

// A Point represents a 2d cartesian coordinate
type Point struct {
	x float64
	y float64
}

// SecondHand is the unit vector of the seonc hand of the analogue clock at time 't'
// represented as a point
func SecondHand(t time.Time) Point {
	return Point{}
}
