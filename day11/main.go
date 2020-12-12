package main

import (
	"advent/day11/seatingchart"
	"fmt"
)

func main() {
	fs := seatingchart.NewFerryChartSimulator("day11/input.txt")
	for i := 0; i < 10; i++ {
		fmt.Printf("Step %d:\n", i)
		fs.PrintCurrentState()
		fs.Step()
		fmt.Println()
	}
	fs.Run()
	fmt.Printf("Occupied seats: %d\n", fs.CountOccupiedSeats())
}
