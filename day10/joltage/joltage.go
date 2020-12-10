package joltage

import (
	"log"
	"strconv"
)

type joltProcessor struct {
	maxExpectedDiff int
	maxJoltage      int
	adapterJoltages map[int]bool
}

func NewJoltProcessor(maxExpectedDiff int) *joltProcessor {
	return &joltProcessor{
		maxExpectedDiff: maxExpectedDiff,
		maxJoltage:      0,
		adapterJoltages: make(map[int]bool),
	}
}

func (p *joltProcessor) ProcessLine(i int, line string) error {
	num, err := strconv.Atoi(line)
	if err != nil {
		return err
	}

	if i == 0 || num > p.maxJoltage {
		p.maxJoltage = num
	}

	p.adapterJoltages[num] = true

	return nil
}

func (p *joltProcessor) GetDiffDistribution() []int {
	diffDistribution := make([]int, p.maxExpectedDiff)
	diffDistribution[p.maxExpectedDiff-1] = 1

	var prev int
	for i := 1; i <= p.maxJoltage; i++ {
		if !p.adapterJoltages[i] {
			continue
		}

		diff := i - prev
		if diff > p.maxExpectedDiff {
			log.Fatalf("Unexpectedly large diff at element %d: %d", i, diff)
		}
		diffDistribution[diff-1]++
		prev = i
	}
	return diffDistribution
}

func (p *joltProcessor) CountArrangements() uint {
	arrangementCounts := make([]uint, p.maxJoltage+1)

	arrangementCounts[0] = 1
	for i := 1; i <= p.maxJoltage; i++ {
		if !p.adapterJoltages[i] {
			continue
		}

		arrangementsUpToN := arrangementCounts[i-1]
		if i-2 >= 0 {
			arrangementsUpToN += arrangementCounts[i-2]
		}
		if i-3 >= 0 {
			arrangementsUpToN += arrangementCounts[i-3]
		}
		arrangementCounts[i] = arrangementsUpToN
	}

	return arrangementCounts[p.maxJoltage]
}
