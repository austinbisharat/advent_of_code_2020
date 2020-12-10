package joltage

import (
	"container/list"
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
	arrangementCounts := list.New()

	arrangementCounts.PushBack(uint(1))
	for i := 1; i <= p.maxJoltage; i++ {
		if !p.adapterJoltages[i] {
			arrangementCounts.PushBack(uint(0))
		} else {
			var sum uint
			for cur := arrangementCounts.Front(); cur != nil; cur = cur.Next() {
				sum += cur.Value.(uint)
			}
			arrangementCounts.PushBack(sum)
		}

		if arrangementCounts.Len() > p.maxExpectedDiff {
			arrangementCounts.Remove(arrangementCounts.Front())
		}
	}

	return arrangementCounts.Back().Value.(uint)
}
