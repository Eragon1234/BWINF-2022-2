package vector

import "math"

func Distance(from, to Coordinate) float64 {
	a := from.X - to.X
	b := from.Y - to.Y
	return math.Sqrt(a*a + b*b)
}
