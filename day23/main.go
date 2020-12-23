package main

import (
	"advent/day23/cups"
	"fmt"
	"strconv"
)

func main() {
	valString := "364297581"
	values := parseValues(valString)
	cr := cups.NewCupRingFromValues(values)
	for i := 0; i < 100; i++ {
		cr.DoRound()
	}
	fmt.Printf("Order after 100: %s\n", cr.ComputeOrderStr())

	values = expandValues(values)
	cr = cups.NewCupRingFromValues(values)
	for i := 0; i < 10000000; i++ {
		cr.DoRound()
	}

	v1, v2 := cr.GetValuesAfterOne()
	fmt.Printf("Next two values after 1: %d * %d = %d\n", v1, v2, v1*v2)
}

func parseValues(valString string) []int {
	var vals []int
	for _, r := range valString {
		val, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}
	return vals
}

func expandValues(values []int) []int {
	newValues := make([]int, len(values), 1000000)
	copy(newValues, values)
	for i := len(values) + 1; i <= 1000000; i++ {
		newValues = append(newValues, i)
	}
	return newValues
}
