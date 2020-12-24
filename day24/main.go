package main

import (
	"advent/common/lineprocessor"
	"advent/day24/hex"
	"fmt"
)

func main() {
	hp := &hexProcessor{board: make(hex.Board)}
	lineprocessor.ProcessLinesInFile("day24/input.txt", hp)
	b := hp.board
	fmt.Printf("count black: %d\n", b.CountBlack())
	for i := 0; i < 100; i++ {
		fmt.Printf("count black after %d steps: %d\n", i, b.CountBlack())
		b.Step()
	}
	fmt.Printf("count black after %d steps: %d\n", 100, b.CountBlack())
}

type hexProcessor struct {
	board hex.Board
}

func (h *hexProcessor) ProcessLine(_ int, line string) error {
	path := hex.NewPathFromStr(line)
	c := hex.Coordinate{}
	c.FollowPath(path)
	h.board.Flip(c)
	return nil
}
