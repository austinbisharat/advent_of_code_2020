package conway

import (
	"fmt"
	"math"
	"strings"
)

const Dimensions = 4

type gridLocation [Dimensions]int

type grid struct {
	activatedLocations map[gridLocation]bool
	minCoord           [Dimensions]int
	maxCoord           [Dimensions]int
}

func newGrid() *grid {
	return &grid{
		activatedLocations: make(map[gridLocation]bool),
		minCoord:           [Dimensions]int{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		maxCoord:           [Dimensions]int{math.MinInt32, math.MinInt32, math.MinInt32},
	}
}

func (g *grid) isActivated(location gridLocation) bool {
	return g.activatedLocations[location]
}

func (g *grid) setActivated(location gridLocation) {
	g.activatedLocations[location] = true

	for i, c := range location {
		if g.minCoord[i] > c {
			g.minCoord[i] = c
		}

		if g.maxCoord[i] < c {
			g.maxCoord[i] = c
		}
	}
}

func (g *grid) unsetActivated(location gridLocation) {
	delete(g.activatedLocations, location)
}

func (g *grid) forEachNeighbor(location gridLocation, visitor gridVisitor) {
	g.forEachNeighborHelper(location, gridLocation{}, 0, visitor)
}

func (g *grid) forEachNeighborHelper(location gridLocation, partialNeighborLocation gridLocation, numCoordSet int, visitor gridVisitor) {
	if numCoordSet == len(location) && location != partialNeighborLocation {
		visitor(g, partialNeighborLocation)
		return
	} else if numCoordSet == len(location) {
		return
	}

	for c := location[numCoordSet] - 1; c <= location[numCoordSet]+1; c++ {
		partialNeighborLocation[numCoordSet] = c
		g.forEachNeighborHelper(location, partialNeighborLocation, numCoordSet+1, visitor)
	}
}

func (g *grid) printGrid() {
	fmt.Println("Grid:")
	for z := g.minCoord[2]; z <= g.maxCoord[2]; z++ {
		fmt.Println(g.sPrintPlane(z))
	}
	fmt.Println()
}

func (g *grid) sPrintPlane(z int) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("z = %d:\n", z))
	for x := g.minCoord[0]; x <= g.maxCoord[0]; x++ {
		for y := g.minCoord[1]; y <= g.maxCoord[1]; y++ {
			if g.isActivated(gridLocation{x, y, z}) {
				builder.WriteString("#")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type gridVisitor func(grid *grid, location gridLocation)

type GridSimulator struct {
	curStep uint64
	curGrid *grid
}

func (s *GridSimulator) Step() {
	nextGrid := newGrid()

	inactiveLocationCounts := make(map[gridLocation]uint8)

	for location := range s.curGrid.activatedLocations {
		var count uint8
		s.curGrid.forEachNeighbor(location, func(_ *grid, neighbor gridLocation) {
			if s.curGrid.isActivated(neighbor) {
				count++
			} else {
				inactiveLocationCounts[neighbor]++

				if inactiveLocationCounts[neighbor] == 3 {
					nextGrid.setActivated(neighbor)
				} else if inactiveLocationCounts[neighbor] > 3 {
					nextGrid.unsetActivated(neighbor)
				}
			}
		})

		if count == 2 || count == 3 {
			nextGrid.setActivated(location)
		}
	}

	s.curStep++
	s.curGrid = nextGrid
}

func (s *GridSimulator) CountActive() int {
	return len(s.curGrid.activatedLocations)
}

func (s *GridSimulator) Print() {
	s.curGrid.printGrid()
}

type GridSimulatorBuilder struct {
	g *grid
}

func NewGridSimulatorBuilder() *GridSimulatorBuilder {
	return &GridSimulatorBuilder{g: newGrid()}
}

func (gsb *GridSimulatorBuilder) ProcessLine(i int, line string) error {
	if gsb.g == nil {
		gsb.g = newGrid()
	}

	for j, r := range line {
		if r == '#' {
			gsb.g.setActivated(gridLocation{i, j})
		}
	}
	return nil
}

func (gsb *GridSimulatorBuilder) Build() *GridSimulator {
	return &GridSimulator{
		curStep: 0,
		curGrid: gsb.g,
	}
}
