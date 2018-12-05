package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"

	"github.com/Gramatiik/go-tsp-aco/aco"
	"github.com/Gramatiik/go-tsp-aco/dataset"
	"github.com/Gramatiik/go-tsp-aco/graph"
	"github.com/Gramatiik/go-tsp-aco/tsplib"
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
	var vertices []*graph.Vertex
	if *filename != "" {
		vertices = tsplib.LoadFromFile(*filename)
	} else {
		// if no file was provided, load the Oliver30 dataset
		vertices = dataset.OLIVER30
	}

	var tspGraph graph.Graph

	// Add vertices to the graph
	for i := 0; i < len(vertices); i++ {
		tspGraph.AddVertex(vertices[i])
	}

	// Create the connections between vertices (complete graph)
	for i := 0; i < len(vertices); i++ {
		for j := 0; j < len(vertices); j++ {
			if vertices[i] != vertices[j] {
				tspGraph.AddEdge(vertices[i], vertices[j])
			}
		}
	}

	// Print a recap of the data
	c := color.New(color.FgGreen, color.Bold)
	d := color.New(color.FgCyan, color.Bold)
	d.Println("\nParameters :")
	d.Println("ANTS\tGENS\tALPHA\tBETA\tEVAP")
	c.Printf("%d\t%d\t%d\t%d\t%.2f\n", *ants, *generations, *alpha, *beta, *evaporationRate)

	fmt.Println("\nNumber of vertices : ", c.Sprintf("%d", tspGraph.GetVerticesCount()))
	fmt.Println("Number of edges : ", c.Sprintf("%d", tspGraph.GetEdgesCount()))
	fmt.Println("")

	tsp := aco.NewTSP(
		&tspGraph,
		*alpha,
		*beta,
		*ants,
		*generations,
		*evaporationRate,
	)

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("cyan")
	s.Start()

	// run the TSP solver and get the best ant
	best := tsp.Run()

	s.FinalMSG = fmt.Sprintf("Best ant did a tour of %s units\n", c.Sprintf("%.3f", best.Evaluate()))
	s.Stop()
}
