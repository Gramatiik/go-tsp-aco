package main

import (
	"fmt"
	"math/rand"
	"time"

	"./aco"
	g "./graph"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var graph g.Graph

	vertices := []g.Vertex{
		g.Vertex{Name: "A", Position: g.Coords{X: 54, Y: 67}},
		g.Vertex{Name: "B", Position: g.Coords{X: 54, Y: 62}},
		g.Vertex{Name: "C", Position: g.Coords{X: 37, Y: 84}},
		g.Vertex{Name: "D", Position: g.Coords{X: 41, Y: 94}},
		g.Vertex{Name: "E", Position: g.Coords{X: 2, Y: 99}},
		g.Vertex{Name: "F", Position: g.Coords{X: 7, Y: 64}},
		g.Vertex{Name: "G", Position: g.Coords{X: 25, Y: 62}},
		g.Vertex{Name: "H", Position: g.Coords{X: 22, Y: 60}},
		g.Vertex{Name: "I", Position: g.Coords{X: 18, Y: 54}},
		g.Vertex{Name: "J", Position: g.Coords{X: 4, Y: 50}},
		g.Vertex{Name: "K", Position: g.Coords{X: 13, Y: 40}},
		g.Vertex{Name: "L", Position: g.Coords{X: 18, Y: 40}},
		g.Vertex{Name: "M", Position: g.Coords{X: 24, Y: 42}},
		g.Vertex{Name: "N", Position: g.Coords{X: 25, Y: 38}},
		g.Vertex{Name: "O", Position: g.Coords{X: 44, Y: 35}},
		g.Vertex{Name: "P", Position: g.Coords{X: 41, Y: 26}},
		g.Vertex{Name: "Q", Position: g.Coords{X: 45, Y: 21}},
		g.Vertex{Name: "R", Position: g.Coords{X: 58, Y: 35}},
		g.Vertex{Name: "S", Position: g.Coords{X: 62, Y: 32}},
		g.Vertex{Name: "T", Position: g.Coords{X: 82, Y: 7}},
		g.Vertex{Name: "U", Position: g.Coords{X: 91, Y: 38}},
		g.Vertex{Name: "V", Position: g.Coords{X: 83, Y: 46}},
		g.Vertex{Name: "W", Position: g.Coords{X: 71, Y: 44}},
		g.Vertex{Name: "X", Position: g.Coords{X: 64, Y: 60}},
		g.Vertex{Name: "X", Position: g.Coords{X: 68, Y: 58}},
		g.Vertex{Name: "Z", Position: g.Coords{X: 83, Y: 69}},
		g.Vertex{Name: "AA", Position: g.Coords{X: 87, Y: 76}},
		g.Vertex{Name: "AB", Position: g.Coords{X: 74, Y: 78}},
		g.Vertex{Name: "AC", Position: g.Coords{X: 71, Y: 71}},
		g.Vertex{Name: "AD", Position: g.Coords{X: 58, Y: 69}},
	}

	// add all vertices
	for i := 0; i < len(vertices); i++ {
		graph.AddVertex(&vertices[i])
	}

	// create connections between them (complete graph)
	for i := 0; i < len(vertices); i++ {
		for j := 0; j < len(vertices); j++ {
			if vertices[i] != vertices[j] {
				graph.AddEdge(&vertices[i], &vertices[j])
			}
		}
	}

	fmt.Println("Number of vertices : ", graph.GetVerticesCount())
	fmt.Println("Number of edges : ", graph.GetEdgesCount())
	fmt.Println("")

	tsp := aco.NewTSP(&graph, 1, 5, 25, 100, 0.5)

	tsp.Run()
}
