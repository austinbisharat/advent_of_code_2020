package main

import (
	"advent/common/lineprocessor"
	"advent/day19/parser"
	"fmt"
)

func main() {
	p1 := parser.NewRecursiveParser("day19/grammar.txt")
	v1 := &lineValidator{p: p1}

	p2 := parser.NewRecursiveParser("day19/grammar2.txt")
	v2 := &lineValidator{p: p2}

	mp := &lineprocessor.LineMultiProcessor{
		Processors: []lineprocessor.LineProcessor{v1, v2},
	}
	lineprocessor.ProcessLinesInFile("day19/input.txt", mp)
	fmt.Printf("Count 1: %d\n", v1.count)
	fmt.Printf("Count 2: %d\n", v2.count)
}

type lineValidator struct {
	p     parser.Parser
	count int
}

func (v *lineValidator) ProcessLine(_ int, line string) error {
	if v.p.IsValid(line) {
		v.count++
	}
	return nil
}
