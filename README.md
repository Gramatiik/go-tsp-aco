# TSP solver using Ant Colony Optimization

This project is a Go implementation of the **Ant System algorithm** proposed by M. Dorigo et al.

The main goal of this project was to produce a functional implementation of this algorithm and to test it against real world data to compare performance and accuracy of the algorithm. It was also a nice way to learn the basics of the Go language !

## Installation

Download the [latest](https://github.com/Gramatiik/go-tsp-aco/releases/latest) version for macOS under the [releases section](https://github.com/Gramatiik/go-tsp-aco/releases)

### Build from source

> To build from source, you will need the [go](https://golang.org/doc/install) compiler and the [vgo](https://github.com/golang/vgo) tool 

Clone this repository and move to the sources directory

```console
$ git clone <repository-url> go-tsp-aco
$ cd go-tsp-aco
```

```console
$ vgo build
```

And _voilà_ ! The compiler generated a binary named `go-tsp-aco` 🚀

## Usage

The following flags are available :
- `-input` : loads the given `.tsp` file.
- `-alpha` : value of the alpha parameter.
- `-beta` : value of the beta parameter.
- `-ants` : number of ants per generations.
- `-generations` : number of generations to simulate.
- `-evaporation` : value for the pheromones evaporation rate.
- `-h or --help` : displays the help message.

## Licence

[MIT](/LICENSE.md)
