package main

import (
	"advent/common/lineprocessor"
	"advent/day17/conway"
	"fmt"
)

func main() {
	b := conway.NewGridSimulatorBuilder()
	lineprocessor.ProcessLinesInFile("day17/input.txt", b)

	gridSimulator := b.Build()
	for i := 0; i < 7; i++ {
		fmt.Printf("After %d steps, %d active\n", i, gridSimulator.CountActive())
		//gridSimulator.Print()
		gridSimulator.Step()
	}
}
