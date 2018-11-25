package tsplib

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"../graph"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// LoadFromFile : load the given tsplib file
func LoadFromFile(filename string) []*graph.Vertex {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	points := []*graph.Vertex{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// next lines are the coords section
		if strings.Contains(scanner.Text(), "NODE_COORD_SECTION") {
			scanner.Scan() // go to node declaration line
			for ; !strings.Contains(scanner.Text(), "EOF"); scanner.Scan() {
				parts := strings.Split(scanner.Text(), " ")

				x, err := strconv.ParseFloat(parts[1], 64)
				check(err)
				y, err := strconv.ParseFloat(parts[2], 64)
				check(err)

				points = append(points, &graph.Vertex{
					Name: string(parts[0]),
					Position: graph.Coords{
						X: x,
						Y: y,
					},
				})
			}
		}
	}

	check(scanner.Err())

	return points
}
