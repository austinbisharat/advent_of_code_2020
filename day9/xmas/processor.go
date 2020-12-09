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
	elements            []int
}

func NewXmasStreamProcessor(k int) *xmasStreamProcessor {
	return &xmasStreamProcessor{
		k:                     k,
		lastKElements:         list.New(),
		lastKElementsCountSet: make(map[int]int),

		firstInvalidElement: nil,
	}
}

func (p *xmasStreamProcessor) ProcessLine(i int, line string) error {
	num, err := strconv.Atoi(line)
	if err != nil {
		return err
	}

	p.elements = append(p.elements, num)

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

func (p *xmasStreamProcessor) GetFirstInvalidNum() int {
	return *p.firstInvalidElement
}

func (p *xmasStreamProcessor) GetMinMaxFromInvalidSumRun() (int, int) {
	low, high := p.getInvalidSumRun()

	var max *int
	var min *int
	for i := low; i <= high; i++ {
		num := p.elements[i]

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

	indexByPrefixSum := map[int]int{}
	var sum int
	for i, element := range p.elements {
		sum += element

		if sum == target && i > 0 {
			return 0, i
		} else if lowIdx, present := indexByPrefixSum[sum-target]; present {
			return lowIdx + 1, i
		}

		indexByPrefixSum[sum] = i
	}

	panic("uhoh")
}
