package main

import (
	"advent/common/lineprocessor"
	"advent/day9/xmas"
	"fmt"
)

func main() {
	processor := xmas.NewXmasStreamProcessor(25)
	lineprocessor.ProcessLinesInFile("day9/input.txt", processor)
	fmt.Printf("First invalid num: %d\n", processor.GetFirstInvalidNum())
	min, max := processor.GetMinMaxFromInvalidSumRun()
	fmt.Printf("Range min max: %d + %d = %d\n", min, max, min+max)
}
