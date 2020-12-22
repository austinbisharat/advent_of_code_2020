package photogrid

import (
	"advent/common/lineprocessor"
	"fmt"
	"regexp"
	"strconv"
)

const (
	tileSize = 10
)

type PhotoTile struct {
	tileID int
	text   [tileSize]string

	borders [2][4]uint16
}

func (tile *PhotoTile) init() {
	for i := 0; i < tileSize; i++ {
		tile.updateBorder(tile.text[0][i], 0, i)
		tile.updateBorder(tile.text[i][tileSize-1], 1, i)
		tile.updateBorder(tile.text[tileSize-1][i], 2, i)
		tile.updateBorder(tile.text[i][0], 3, i)
	}
}

func (tile *PhotoTile) updateBorder(r uint8, borderNum, index int) {
	if r == '#' {
		tile.borders[0][borderNum] |= 1 << (tileSize - index - 1)
		tile.borders[1][borderNum] |= 1 << index
	}
}

type tileReader struct {
	curTile *PhotoTile
	tiles   []PhotoTile
}

var tileRegex = regexp.MustCompile("^Tile ([0-9]+):$")

func (t *tileReader) ProcessLine(i int, line string) error {
	i %= 12
	if i == 0 {
		submatches := tileRegex.FindStringSubmatch(line)
		if len(submatches) != 2 {
			return fmt.Errorf("invalid title line '%s'", line)
		}
		tileID, _ := strconv.Atoi(submatches[1])
		t.curTile = &PhotoTile{
			tileID:  tileID,
			text:    [10]string{},
			borders: [2][4]uint16{},
		}
	} else if i < 10 {
		t.curTile.text[i-1] = line
	} else if i == 10 {
		t.curTile.text[i-1] = line
		t.curTile.init()
		t.tiles = append(t.tiles, *t.curTile)
	}
	return nil
}

func ReadTiles(path string) []PhotoTile {
	r := &tileReader{}
	lineprocessor.ProcessLinesInFile(path, r)
	return r.tiles
}
