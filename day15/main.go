package main

import "fmt"

func main() {
	inputArr := []int{2, 0, 1, 9, 5, 19}
	fmt.Printf("Seq %v element %d is %d\n", inputArr, 2020, computeNumInSequence(inputArr, 2020))
	fmt.Printf("Seq %v element %d is %d\n", inputArr, 30000000, computeNumInSequence(inputArr, 30000000))
}

func computeNumInSequence(initialSequence []int, n int) int {
	prevIndexByNum := make(map[int]int)
	prev := 0
	for i := 0; i < n; i++ {
		var cur int
		if i < len(initialSequence) {
			cur = initialSequence[i]
		} else if prevIdx, exists := prevIndexByNum[prev]; exists {
			cur = i - prevIdx - 1
		} else {
			cur = 0
		}
		prevIndexByNum[prev] = i - 1
		prev = cur
	}
	return prev
}
