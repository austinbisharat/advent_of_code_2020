package main

import (
	"advent/common/lineprocessor"
	"advent/day10/joltage"
	"fmt"
)

func main() {
	jp := joltage.NewJoltProcessor(3)
	lineprocessor.ProcessLinesInFile("day10/input.txt", jp)
	diffs := jp.GetDiffDistribution()
	fmt.Printf("Count 1s: %d, count 3s: %d, product: %d\n", diffs[0], diffs[2], diffs[0]*diffs[2])
	fmt.Printf("arrangments: %d\n", jp.CountArrangements())
}
