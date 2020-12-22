package main

import (
	"advent/day20/photogrid"
	"fmt"
)

func main() {

	tiles := photogrid.ReadTiles("day20/input.txt")
	seaMonster := photogrid.NewPatternFromFile("day20/pattern.txt")
	pointPatter := photogrid.NewPointPattern()

	solver := photogrid.NewPhotoGridSolver(tiles)
	solutions := solver.Solve()
	fmt.Printf("Has %d solutions\n", len(solutions))
	solution := solutions[0]

	fmt.Printf("Corner product: %d\n", solution.GetCornerProduct())
	photo := photogrid.NewPhoto(solution)

	pointPatternFinder := photogrid.NewPatternFinder(pointPatter, &photo)
	numPoints := pointPatternFinder.CountPattern()

	seaMonsterFinder := photogrid.NewPatternFinder(seaMonster, &photo)
	n := seaMonsterFinder.CountPattern()
	fmt.Printf("Num points %d, num sea monsters %d, score %d\n", numPoints, n, numPoints-n*seaMonster.Size())

}
