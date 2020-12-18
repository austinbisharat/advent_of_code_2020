package main

import (
	"advent/common/lineprocessor"
	"advent/day18/math"
	"fmt"
)

func main() {
	p := &MathProcessor{}
	lineprocessor.ProcessLinesInFile("day18/input.txt", p)
	fmt.Printf("Sum 1: %d\n", p.part1)
	fmt.Printf("Sum 2: %d\n", p.part2)
}

type MathProcessor struct {
	part1 int
	part2 int
}

func (p *MathProcessor) ProcessLine(_ int, line string) error {
	value1, value2, err := math.EvaluateExpression(line)
	p.part1 += value1
	p.part2 += value2
	return err
}
