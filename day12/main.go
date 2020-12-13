package main

import (
	"advent/common/lineprocessor"
	"advent/day12/navigation"
	"fmt"
)

func main() {
	processor := &FerryProcessor{
		location: navigation.NewFerryLocation(),
	}
	lineprocessor.ProcessLinesInFile("day12/input.txt", processor)
	fmt.Printf("Distance: %d\n", processor.location.DistanceFromOrigin())
}

type FerryProcessor struct {
	location *navigation.FerryLocation
}

func (p *FerryProcessor) ProcessLine(_ int, line string) error {
	var instruction navigation.FerryInstruction
	err := instruction.Parse(line)
	p.location.ApplyInstruction(instruction)
	return err
}
