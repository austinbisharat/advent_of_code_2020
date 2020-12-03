package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	candidateSlopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1

	for _, treeCount := range getCountsForCandidateSlopes(candidateSlopes) {
		product *= treeCount
	}

	fmt.Printf("Product: %v\n", product)
}

type Slope struct {
	right int
	down int
}

func getCountsForCandidateSlopes(slopes []Slope) []int {
	tobogganProcessors := make([]LineProcessor, 0, len(slopes))

	for _, slope := range slopes {
		tobogganProcessors = append(tobogganProcessors, NewTobogginProcessor(slope))
	}

	processLinesInFile("day3/input.txt", &LineMultiProcessor{tobogganProcessors})

	treeCounts := make([]int, 0, len(slopes))
	for _, processor := range tobogganProcessors {
		treeCounts = append(treeCounts, processor.(*TobogganProcessor).treeCount)
	}
	return treeCounts
}


type LineProcessor interface {
	ProcessLine(int, []rune) error
}

func processLinesInFile(path string, processor LineProcessor) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		processor.ProcessLine(i, []rune(line))
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type LineMultiProcessor struct {
	processors []LineProcessor
}

func (m *LineMultiProcessor) ProcessLine(i int, line []rune) error {
	for _, processor := range m.processors {
		err := processor.ProcessLine(i, line)
		if err != nil {
			return err
		}
	}
	return nil
}

type TobogganProcessor struct {
	slope Slope

	lineLength int
	verticalPos int
	horizontalPos int

	treeCount int
}

func NewTobogginProcessor(slope Slope) *TobogganProcessor {
	return &TobogganProcessor{
		slope: slope,

		lineLength: -1,

		verticalPos: 0,
		horizontalPos: 0,

		treeCount: 0,
	}
}

func (t *TobogganProcessor) ProcessLine(i int, line []rune) error {
	if t.lineLength != -1 && len(line) != t.lineLength {
		return errors.New("variable line lengths")
	}
	t.lineLength = len(line)

	if i != t.verticalPos {
		return nil
	}

	if line[t.horizontalPos % t.lineLength] == '#' {
		t.treeCount++
	}

	t.horizontalPos += t.slope.right
	t.verticalPos += t.slope.down

	return nil
}

