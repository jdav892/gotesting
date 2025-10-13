package clockface

import (
	"math"
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
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}
	p = Point{p.X, -p.Y}
	p = Point{p.X + 150, p.Y + 150}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
