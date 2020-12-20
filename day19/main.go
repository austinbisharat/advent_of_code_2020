package main

import (
	"advent/common/lineprocessor"
	"advent/day19/parser"
	"fmt"
)

func main() {
	p := parser.NewBruteForceParser("day19/grammar.txt")
	v := &lineValidator{p: p}
	lineprocessor.ProcessLinesInFile("day19/input.txt", v)
	fmt.Printf("Count: %d\n", v.count)
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
