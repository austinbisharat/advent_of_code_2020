package xmas

import (
	"container/list"
	"strconv"
)

type xmasStreamProcessor struct {
	k                     int
	lastKElements         *list.List
	lastKElementsCountSet map[int]int

	firstInvalidElement *int
	prefixSums          []int
}

func NewXmasStreamProcessor(k int) *xmasStreamProcessor {
	return &xmasStreamProcessor{
		k:                     k,
		lastKElements:         list.New(),
		lastKElementsCountSet: make(map[int]int),

		firstInvalidElement: nil,
		prefixSums:          []int{0},
	}
}

func (p *xmasStreamProcessor) ProcessLine(i int, line string) error {
	num, err := strconv.Atoi(line)
	if err != nil {
		return err
	}

	p.updatePrefixSums(num)

	if p.lastKElements.Len() < p.k {
		p.lastKElements.PushBack(num)
		p.lastKElementsCountSet[num]++
		return nil
	}

	p.checkValidity(num)

	p.updateLastKElements(num)

	return nil
}

func (p *xmasStreamProcessor) updateLastKElements(num int) {
	elementToRemove := p.lastKElements.Front()
	numToRemove := elementToRemove.Value.(int)
	p.lastKElements.Remove(elementToRemove)
	if p.lastKElementsCountSet[numToRemove] > 1 {
		p.lastKElementsCountSet[numToRemove]--
	} else {
		delete(p.lastKElementsCountSet, numToRemove)
	}

	p.lastKElements.PushBack(num)
	p.lastKElementsCountSet[num]++
}

func (p *xmasStreamProcessor) checkValidity(num int) {
	var isValid bool
	for element, count := range p.lastKElementsCountSet {
		targetNum := num - element

		isValid = (targetNum == element && count >= 2) || (p.lastKElementsCountSet[targetNum] > 0)
		if isValid {
			break
		}
	}

	if !isValid && p.firstInvalidElement == nil {
		p.firstInvalidElement = &num
	}
}

func (p *xmasStreamProcessor) updatePrefixSums(num int) {
	p.prefixSums = append(p.prefixSums, num+p.prefixSums[len(p.prefixSums)-1])
}

func (p *xmasStreamProcessor) GetFirstInvalidNum() int {
	return *p.firstInvalidElement
}

func (p *xmasStreamProcessor) GetMinMaxFromInvalidSumRun() (int, int) {
	low, high := p.getInvalidSumRun()

	var max *int
	var min *int
	for i := low; i < high; i++ {
		num := p.prefixSums[i+1] - p.prefixSums[i]

		if min == nil || *min < num {
			min = &num
		}

		if max == nil || *max > num {
			max = &num
		}
	}
	return *min, *max
}

func (p *xmasStreamProcessor) getInvalidSumRun() (low, high int) {
	target := *p.firstInvalidElement

	for low := 0; low < len(p.prefixSums)-2; low++ {
		lowHigh := low + 2
		highHigh := len(p.prefixSums)
		for lowHigh <= highHigh {
			high = ((highHigh + lowHigh - 1) / 2) + 1
			rangeSum := p.computeRangeSum(low, high)
			if rangeSum > target {
				highHigh = high - 1
			} else if rangeSum < target {
				lowHigh = high + 1
			} else {
				return low, high
			}
		}
	}

	panic("uhoh")
}

// computes sum of the range starting at index low inclusive to high exclusive
func (p *xmasStreamProcessor) computeRangeSum(low, high int) int {
	return p.prefixSums[high] - p.prefixSums[low]
}
