package graph

import "math"

// Edge : represents an edge on the graph
type Edge struct {
	First      *Vertex
	Second     *Vertex
	Pheromones float64
	Weight     float64
}

// GetOppositeEnd : Return the opposite end vertex if both are connected
// returns nil otherwise
func (e *Edge) GetOppositeEnd(v *Vertex) *Vertex {
	if e.First == v {
		return e.Second
	} else if e.Second == v {
		return e.First
	} else {
		return nil
	}
}

// Hash : Identifying hash of the edge
func (e *Edge) Hash() float64 {
	return math.Sqrt(e.First.Hash()) + math.Sqrt(e.Second.Hash())
}
