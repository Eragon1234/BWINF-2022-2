package vector

import (
	"fmt"
)

type Coordinate struct {
	X, Y float64
}

// ParseCoordinate parses a coordinate with x and y values separated by a space.
func ParseCoordinate(s string) (Coordinate, error) {
	var c Coordinate
	_, err := fmt.Sscanf(s, "%f %f", &c.X, &c.Y)
	return c, err
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%.6f %.6f", c.X, c.Y)
}
