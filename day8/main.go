package main

import (
	"advent/day8/program"
	"fmt"
)

func main() {
	p := program.LoadProgram("day8/input.txt")
	r := program.NewRuntime(p)
	r.RunUntilInfiniteLoop()
	fmt.Printf("accumulator before infinite loop: %d\n", r.GetAccumulator())

	// N^2 runtime
	r = program.NewRuntime(p)
	acc := r.RunSelfCorrecting()
	fmt.Printf("accumulator after terminal: %d\n", acc)

	// linear runtime
	p.RemoveInfiniteLoop()
	r = program.NewRuntime(p)
	r.RunUntilInfiniteLoop()
}
