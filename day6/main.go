package main

import (
	"advent/common/lineprocessor"
	"fmt"
)

func main() {
	processor := &processor{}
	lineprocessor.ProcessLinesInFile("day6/input.txt", processor)
	fmt.Printf("sum ticket: %d\n", processor.GetCount())
}

type processor struct {
	answerSet *letterSet
	sum uint
}

func (p *processor) ProcessLine(_ int, line string) error {
	if line != "" {
		ls := &letterSet{}

		for _, r := range line {
			ls.Add(r)
		}

		if p.answerSet == nil {
			p.answerSet = ls
		} else {
			p.answerSet.Retain(ls)
		}
	} else if p.answerSet != nil {
		p.sum += p.answerSet.GetCount()
		p.answerSet = nil
	}
	return nil
}

func (p *processor) GetCount() uint {
	if p.answerSet != nil {
		p.sum += p.answerSet.GetCount()
	}
	return p.sum
}

type letterSet struct {
	set uint32
	count uint
}

func (l *letterSet) Add(r rune) {
	idx := int(r) - int('a')
	if (l.set >> idx) & 1 == 0 {
		l.count++
		l.set |= 1 << idx
	}
}

func (l *letterSet) GetCount() uint {
	return l.count
}

func (l *letterSet) Retain(other *letterSet) {
	before := l.set
	l.set &= other.set

	bitsChanged := before ^ l.set
	for bitsChanged != 0 {
		bitsChanged &= bitsChanged - 1
		l.count--
	}
}
