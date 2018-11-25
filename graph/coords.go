package graph

import "math"

// Coords : Holds X/Y 2d coordinates
type Coords struct {
	X float64
	Y float64
}

// DistanceTo : Calculates the distance between two coordinates
func (c *Coords) DistanceTo(other *Coords) float64 {
	dx := c.X - other.X
	dy := c.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}
