package photogrid

import (
	"advent/common/lineprocessor"
	"fmt"
	"os"
)

const (
	photoSize = photoGridSize * (tileSize - 2)
)

type photo struct {
	image [photoSize][photoSize]bool

	orientation orientation
}

func (p *photo) set(row, col int, value bool) {
	p.image[row][col] = value
}

func (p *photo) get(row, col int) bool {
	row, col = p.orientation.translate(row, col)
	return p.image[row][col]
}

func (p *photo) Print() {
	for i := range p.image {
		for j := range p.image[i] {
			if p.image[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (p *photo) PrintFile(file string) {
	f, _ := os.Create(file)
	defer f.Close()
	for flipped := 0; flipped < 2; flipped++ {
		for rotation := uint8(0); rotation < 4; rotation++ {
			p.orientation = orientation{
				flipped:   flipped == 1,
				rotations: rotation,
				size:      photoSize,
			}

			for i := range p.image {
				for j := range p.image[i] {
					if p.get(i, j) {
						f.WriteString("#")
					} else {
						f.WriteString(".")
					}
				}
				f.WriteString("\n")
			}

			f.WriteString("\n\n")
		}
	}

}

func (p *photo) SetOrientation(orientation orientation) {
	p.orientation = orientation
}

func NewPhoto(g Grid) photo {
	p := photo{}
	for row := range g {
		for col := range g[row] {
			placement := g[row][col]

			rs := (tileSize-2)*row - 1
			cs := (tileSize-2)*col - 1
			for tileRow := 1; tileRow < tileSize-1; tileRow++ {
				for tileCol := 1; tileCol < tileSize-1; tileCol++ {
					val := placement.getValue(tileRow, tileCol)
					p.set(rs+tileRow, cs+tileCol, val)
				}
			}
		}
	}
	return p
}

func (p *photo) IsEquivalent(other *photo) bool {
	for flipped := 0; flipped < 2; flipped++ {
		for rotations := uint8(0); rotations < 4; rotations++ {
			p.orientation = orientation{
				flipped:   flipped == 1,
				rotations: rotations,
				size:      photoSize,
			}

			if p.matches(other) {
				return true
			}
		}
	}
	return false
}

func (p *photo) matches(other *photo) bool {
	for row := range p.image {
		for col := range p.image[row] {
			if p.get(row, col) != other.get(row, col) {
				return false
			}
		}
	}
	return true
}

func (p *photo) ProcessLine(row int, line string) error {
	for col := range line {
		p.set(row, col, line[col] == '#')
	}
	return nil
}

func NewPhotoFromFile(file string) *photo {
	p := &photo{}
	lineprocessor.ProcessLinesInFile(file, p)
	return p
}
