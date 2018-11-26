package graph

import (
	"math/rand"
	"sync"
)

// Graph : Represents an non oriented graph with vertices and edges
type Graph struct {
	Name     string
	Vertices map[float64]*Vertex
	Edges    map[float64]*Edge
	mut      sync.RWMutex
}

// AddVertex : Add a vertex to the graph
func (g *Graph) AddVertex(v *Vertex) {
	g.mut.Lock()

	// initialize the map the first time it's used
	if g.Vertices == nil {
		g.Vertices = make(map[float64]*Vertex)
	}

	// if the vertex is not already present, add it
	if _, present := g.Vertices[v.Hash()]; !present {
		g.Vertices[v.Hash()] = v
	}
	g.mut.Unlock()
}

// AddEdge : Add an edge between the two vertices
func (g *Graph) AddEdge(v1, v2 *Vertex) *Edge {
	g.mut.Lock()
	e := Edge{First: v1, Second: v2, Pheromones: 0.1}
	if g.Edges == nil {
		g.Edges = make(map[float64]*Edge)
	}

	// add the edge is it's ot already present
	if _, present := g.Edges[e.Hash()]; !present {
		g.Edges[e.Hash()] = &e
	}
	g.mut.Unlock()

	return &e
}

// IsEmpty : Returns true if the graph has no Vertices, returns false otherwise
func (g *Graph) IsEmpty() bool {
	return len(g.Vertices) == 0
}

// GetVerticesCount : total number of vertices in the graph
func (g *Graph) GetVerticesCount() int {
	return len(g.Vertices)
}

// GetEdgesCount : total number of edges in the graph
func (g *Graph) GetEdgesCount() int {
	return len(g.Edges)
}

// GetRandomVertex : return a random vertex from the vertices of the graph
func (g *Graph) GetRandomVertex() *Vertex {
	rnd := rand.Intn(len(g.Vertices))
	for _, v := range g.Vertices {
		if rnd == 0 {
			return v
		}
		rnd--
	}
	return nil
}

// GetEdgesForVertex : returns the edges connected to the vertex
func (g *Graph) GetEdgesForVertex(v *Vertex) []*Edge {
	ret := []*Edge{}

	for _, e := range g.Edges {
		if e.First == v || e.Second == v {
			ret = append(ret, e)
		}
	}

	return ret
}

// GetEdgeBetweenVertices : return the edge between the given vertices, or nil if it's not present
func (g *Graph) GetEdgeBetweenVertices(v1, v2 *Vertex) *Edge {
	g.mut.Lock()
	e := Edge{First: v1, Second: v2}

	if _, present := g.Edges[e.Hash()]; !present {
		return nil
	}

	defer g.mut.Unlock()
	return g.Edges[e.Hash()]
}
