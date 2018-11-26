package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"./aco"
	"./dataset"
	g "./graph"
	"./tsplib"
)

func main() {
	// seed the random generator (needed for random ant placement)
	rand.Seed(time.Now().UTC().UnixNano())

	// parse the command line aruments
	alpha := flag.Uint("alpha", 1, "Alpha value")
	beta := flag.Uint("beta", 5, "Beta value")
	ants := flag.Uint("ants", 35, "Number of ants per generation")
	generations := flag.Uint("generations", 25, "Number of generations")
	evaporationRate := flag.Float64("evaportation", 0.5, "Evaporation rate of pheromones")
	filename := flag.String("input", "", "File input in TSPLIB format (.tsp),\nuses the Oliver30 data set if not specified")
	flag.Parse()

	// load the dataset (from file or the default one)
	var vertices []*g.Vertex
	if *filename != "" {
		vertices = tsplib.LoadFromFile(*filename)
	} else {
		// if no file was provided, load the Oliver30 dataset
		vertices = dataset.OLIVER30
	}

	var graph g.Graph

	// Add vertices to the graph
	for i := 0; i < len(vertices); i++ {
		graph.AddVertex(vertices[i])
	}

	// Create the connections between vertices (complete graph)
	for i := 0; i < len(vertices); i++ {
		for j := 0; j < len(vertices); j++ {
			if vertices[i] != vertices[j] {
				graph.AddEdge(vertices[i], vertices[j])
			}
		}
	}

	// Print a recap of the data
	fmt.Println("Parameters :")
	fmt.Println("\t- ANTS : ", *ants)
	fmt.Println("\t- GENERATIONS : ", *generations)
	fmt.Println("\t- ALPHA : ", *alpha)
	fmt.Println("\t- BETA : ", *beta)
	fmt.Println("\t- EVAPORATION RATE : ", *evaporationRate)

	fmt.Println("\nNumber of vertices : ", graph.GetVerticesCount())
	fmt.Println("Number of edges : ", graph.GetEdgesCount())
	fmt.Println("")

	tsp := aco.NewTSP(
		&graph,
		*alpha,
		*beta,
		*ants,
		*generations,
		*evaporationRate,
	)

	tsp.Run()
}
