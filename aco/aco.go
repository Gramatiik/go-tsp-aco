package aco

import (
	"fmt"
	"math"
	"math/rand"

	g "github.com/Gramatiik/go-tsp-aco/graph"
)

// TSP : the TSP problem solver using AS algorithm
type TSP struct {
	// internal data
	graph *g.Graph
	ants  []*Ant

	// parameters
	alpha               uint
	beta                uint
	evaporationRate     float64
	numberOfAnts        uint
	numberOfGenerations uint
}

// NewTSP : create a new TSP problem solver
func NewTSP(graph *g.Graph, alpha, beta, ants, generations uint, evaporationRate float64) *TSP {
	var tsp TSP

	tsp.ants = []*Ant{}
	tsp.graph = graph

	tsp.alpha = alpha
	tsp.beta = beta
	tsp.evaporationRate = evaporationRate

	tsp.numberOfAnts = ants
	tsp.numberOfGenerations = generations

	return &tsp
}

// Run : start the TSP solving with the given parameters
func (tsp *TSP) Run() {
	var bestAnt *Ant
	for i := 0; i < int(tsp.numberOfGenerations); i++ {
		tsp.createAnts()

		bestAntOfGeneration := tsp.updateAntsPositions()
		tsp.evaporatePheromones()
		tsp.updatePheromones()

		if bestAnt == nil || bestAnt.Evaluate() > bestAntOfGeneration.Evaluate() {
			bestAnt = bestAntOfGeneration
		}
	}

	fmt.Println("Best ant did a tour of length : ", bestAnt.Evaluate())
}

func (tsp *TSP) createAnts() {
	for i := 0; i < int(tsp.numberOfAnts); i++ {
		tsp.ants = append(tsp.ants, NewAnt(tsp.graph, tsp.alpha, tsp.beta))
	}
}

func (tsp *TSP) updateAntsPositions() *Ant {
	var bestAnt *Ant
	for i := 0; i < len(tsp.ants); i++ {
		for !tsp.ants[i].IsTravelFinished() {
			tsp.ants[i].Travel()
		}

		// best ant is the one with lowest eval
		if bestAnt == nil || bestAnt.Evaluate() > tsp.ants[i].Evaluate() {
			bestAnt = tsp.ants[i]
		}
	}

	return bestAnt
}

func (tsp *TSP) evaporatePheromones() {
	for _, edge := range tsp.graph.Edges {
		edge.Pheromones *= (1 - float64(tsp.evaporationRate))
	}
}

func (tsp *TSP) updatePheromones() {
	for _, ant := range tsp.ants {
		tour := ant.GetTour()
		tourLength := ant.Evaluate()

		for j := 1; j < len(tour); j++ {
			edge := tsp.graph.GetEdgeBetweenVertices(tour[j-1], tour[j])

			if edge != nil {
				edge.Pheromones += 1.0 / tourLength
			}
		}
	}
}

// Ant : the ant agent
type Ant struct {
	graph         *g.Graph
	currentVertex *g.Vertex
	alpha         uint
	beta          uint

	visitedVertices map[float64]*g.Vertex
	tour            []*g.Vertex
}

// NewAnt : create a new ant
func NewAnt(graph *g.Graph, alpha, beta uint) *Ant {
	var ant Ant

	ant.graph = graph
	ant.alpha = alpha
	ant.beta = beta

	// place the ant at a random position on the graph
	initialVertex := graph.GetRandomVertex()
	ant.currentVertex = initialVertex
	ant.visitedVertices = make(map[float64]*g.Vertex)
	ant.visitedVertices[initialVertex.Hash()] = initialVertex
	ant.tour = []*g.Vertex{initialVertex}

	return &ant
}

// Travel : Make the ant travel to the next vertex on the graph
func (a *Ant) Travel() {
	if a.IsTravelFinished() {
		return
	}

	// if the ant visited all the vertices, append the initial vertex to finish the tour
	if len(a.graph.Vertices) == len(a.tour) {
		a.tour = append(a.tour, a.tour[0])
		return
	}

	nextVertex := a.nextVertex()
	a.visitedVertices[nextVertex.Hash()] = nextVertex
	a.tour = append(a.tour, nextVertex)
	a.currentVertex = nextVertex
}

// IsTravelFinished : Tells wether the ant visited the whole graph or not
func (a *Ant) IsTravelFinished() bool {
	return len(a.graph.Vertices)+1 == len(a.tour)
}

// Evaluate : get the score of the tour (it's distance)
func (a *Ant) Evaluate() float64 {
	score := 0.0
	for i := 1; i < len(a.tour); i++ {
		score += a.tour[i].Position.DistanceTo(&a.tour[i-1].Position)
	}
	return score
}

// GetTour : returns the array of visited vertices
func (a *Ant) GetTour() []*g.Vertex {
	if !a.IsTravelFinished() {
		fmt.Println("WARN : The returned tour is incomplete")
	}
	return a.tour
}

// get the next vertex to move to
func (a *Ant) nextVertex() *g.Vertex {
	probs, edges := a.probabilities()
	r := rand.Float64()

	for i := 0; i < len(probs); i++ {
		if r <= probs[i] {
			return edges[i].GetOppositeEnd(a.currentVertex)
		}
	}

	fmt.Println("Unable to choose a next vertex...")
	return nil
}

func (a *Ant) probabilities() ([]float64, []*g.Edge) {
	connectedEdges := a.graph.GetEdgesForVertex(a.currentVertex)
	probabilities := []float64{}
	allowedEdges := []*g.Edge{}

	// if the connected vertex has not been visited, add it to the available vertices list
	for i := 0; i < len(connectedEdges); i++ {
		if _, present := a.visitedVertices[connectedEdges[i].GetOppositeEnd(a.currentVertex).Hash()]; !present {
			allowedEdges = append(allowedEdges, connectedEdges[i])
		}
	}

	denominator := a.denominator(allowedEdges)

	for i := 0; i < len(allowedEdges); i++ {
		probability := a.desirability(allowedEdges[i]) / denominator
		if len(probabilities) != 0 {
			probability += probabilities[len(probabilities)-1]
		}

		probabilities = append(probabilities, probability)
	}

	return probabilities, allowedEdges
}

// calculate the sum of desirability for the given edges
func (a *Ant) denominator(edges []*g.Edge) float64 {
	denominator := 0.0
	for i := 0; i < len(edges); i++ {
		denominator += a.desirability(edges[i])
	}
	return denominator
}

// calculate the desirability for the given edge
func (a *Ant) desirability(edge *g.Edge) float64 {
	pheromone := math.Pow(edge.Pheromones, float64(a.alpha))
	distance := edge.First.Position.DistanceTo(&edge.Second.Position)
	return pheromone * math.Pow(1/distance, float64(a.beta))
}
