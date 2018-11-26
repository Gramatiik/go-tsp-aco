package aco

import (
	"fmt"

	g "../graph"
)

// TSP : the TSP problem solver
type TSP struct {
	graph           *g.Graph
	alpha           uint
	beta            uint
	evaporationRate float64
	ants            uint
	generations     uint
}

// NewTSP : create a new TSP problem solver using ACO
func NewTSP(graph *g.Graph, alpha, beta, ants, generations uint, evaporationRate float64) *TSP {
	var tsp TSP
	tsp.graph = graph

	tsp.alpha = alpha
	tsp.beta = beta
	tsp.evaporationRate = evaporationRate

	tsp.ants = ants
	tsp.generations = generations

	return &tsp
}

// Run : start the TSP solving with the given parameters
func (tsp *TSP) Run() {
	var bestAnt *Ant
	for i := 0; i < int(tsp.generations); i++ {
		ants := tsp.createAnts(tsp.ants)

		bestAntOfGeneration := tsp.updateAntsPositions(ants)
		tsp.updatePheromones(ants)

		if bestAnt == nil || bestAnt.Evaluate() > bestAntOfGeneration.Evaluate() {
			bestAnt = bestAntOfGeneration
		}
	}

	fmt.Println("Best ant did a tour of length : ", bestAnt.Evaluate())
}

func (tsp *TSP) createAnts(n uint) []*Ant {
	ants := []*Ant{}
	for i := 0; i < int(n); i++ {
		ants = append(ants, NewAnt(tsp.graph, tsp.alpha, tsp.beta))
	}
	return ants
}

func (tsp *TSP) updateAntsPositions(ants []*Ant) *Ant {
	var bestAnt *Ant
	for i := 0; i < len(ants); i++ {
		for !ants[i].IsTravelFinished() {
			ants[i].Travel()
		}

		// best ant is the one with lowest eval
		if bestAnt == nil || bestAnt.Evaluate() > ants[i].Evaluate() {
			bestAnt = ants[i]
		}
	}

	return bestAnt
}

func (tsp *TSP) updatePheromones(ants []*Ant) {
	for i := 0; i < len(ants); i++ {
		ant := ants[i]

		tour := ant.GetTour()
		tourLength := ant.Evaluate()

		for j := 1; j < len(tour); j++ {
			edge := tsp.graph.GetEdgeBetweenVertices(tour[j-1], tour[j])

			if edge != nil {
				edge.Pheromones = (1-float64(tsp.evaporationRate))*edge.Pheromones + 1.0/tourLength
			}
		}
	}
}
