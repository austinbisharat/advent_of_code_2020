package photogrid

import "advent/common/lineprocessor"

type point struct {
	row int
	col int
}

type pattern struct {
	relativeOffsets map[point]bool
	colMax          int
	rowMax          int
}

func (p *pattern) ProcessLine(row int, line string) error {
	for col := range line {
		if line[col] == '#' {
			p.relativeOffsets[point{
				row: row,
				col: col,
			}] = true

			if col > p.colMax {
				p.colMax = col
			}
			if row > p.rowMax {
				p.rowMax = row
			}
		}
	}
	return nil
}

func NewPatternFromFile(file string) *pattern {
	p := pattern{
		relativeOffsets: make(map[point]bool),
	}
	lineprocessor.ProcessLinesInFile(file, &p)
	return &p
}

func NewPointPattern() *pattern {
	p := pattern{
		relativeOffsets: map[point]bool{point{
			row: 0,
			col: 0,
		}: true},
		colMax: 0,
		rowMax: 0,
	}
	return &p
}

func (p *pattern) Size() int {
	return len(p.relativeOffsets)
}

type patternFinder struct {
	pattern *pattern
	photo   *photo
}

func NewPatternFinder(pattern *pattern, photo *photo) *patternFinder {
	return &patternFinder{
		pattern: pattern,
		photo:   photo,
	}
}

func (finder *patternFinder) CountPattern() int {
	max := -1
	for flipped := 0; flipped < 2; flipped++ {
		for rotations := uint8(0); rotations < 4; rotations++ {
			finder.photo.SetOrientation(orientation{
				flipped:   flipped == 1,
				rotations: rotations,
				size:      photoSize,
			})
			c := finder.countPattern()
			if c > max {
				max = c
			}
		}
	}
	return max
}

func (finder *patternFinder) countPattern() int {
	var count int
	for row := range finder.photo.image {
		for col := range finder.photo.image {
			if finder.isPatternAt(row, col) {
				count++
			}
		}
	}
	return count
}

func (finder *patternFinder) isPatternAt(row int, col int) bool {
	if row+finder.pattern.rowMax >= photoSize || col+finder.pattern.colMax >= photoSize {
		return false
	}

	for offset := range finder.pattern.relativeOffsets {
		if !finder.photo.get(row+offset.row, col+offset.col) {
			return false
		}
	}

	return true
}
