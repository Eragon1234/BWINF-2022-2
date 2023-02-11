package vector

import (
	"math"
)

func Angle(from, to Coordinate) float64 {
	// calculate the angle between the line of the two points and a horizontal line
	a := to.X - from.X
	b := to.Y - from.Y
	t := math.Atan2(a, b)
	degrees := ToDegrees(t)
	return degrees
}

func ToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
