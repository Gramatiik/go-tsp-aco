package tsplib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../graph"
)

const NAME = "NAME"
const COMMENT = "COMMENT"
const NODE_COORD_SECTION = "NODE_COORD_SECTION"

// Basic error handler
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
		line := scanner.Text()

		switch {
		case strings.Contains(line, NODE_COORD_SECTION): // read coords section
			fmt.Println("parsing NODE_COORD_SECTION...")
			scanner.Scan()
			points = readNodeCoordSection(scanner)
		default:
			fmt.Println("skipped unsupported tsplib section...")
		}
	}

	check(scanner.Err())

	return points
}

// Read the coords NODE_COORD_SECTION part of the file
// The scanner pointer MUST be placed at the first line of nodes coords
// returns an array of parsed vertices
func readNodeCoordSection(scanner *bufio.Scanner) []*graph.Vertex {
	points := []*graph.Vertex{}

	for {
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

		// stop parsing if end of file or EOF marker encountered
		if ok := scanner.Scan(); !ok || strings.Contains(scanner.Text(), "EOF") {
			break
		}
	}

	return points
}
