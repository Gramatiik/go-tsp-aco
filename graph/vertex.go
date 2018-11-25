package graph

import "fmt"

// Vertex represents a node element in a graph
type Vertex struct {
	Name     string
	Position Coords
}

// String : returns a string representation of the vertex
func (v *Vertex) String() string {
	return fmt.Sprintf("%v", v.Name)
}

// Hash : hashcode for the vertex based on its coordinates
func (v *Vertex) Hash() float64 {
	return 31*v.Position.X + v.Position.Y
}
