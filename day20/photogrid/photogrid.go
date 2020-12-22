package photogrid

const (
	photoGridSize = 12
)

type photoTilePlacement struct {
	tile        *PhotoTile
	orientation orientation
}

// TODO: Definitely a way to clean up the redundancy in these methods, but I have
//  spent too much time looking at rotations/flips to care
func (p *photoTilePlacement) getTopBorder() uint16 {
	borderNum := p.orientation.rotations % 4
	isBorderFlipped := borderNum > 1
	if p.orientation.flipped {
		borderNum = (8 - borderNum) % 4
		isBorderFlipped = borderNum < 2
	}
	if isBorderFlipped {
		return p.tile.borders[1][borderNum]
	} else {
		return p.tile.borders[0][borderNum]
	}
}

func (p *photoTilePlacement) getRightBorder() uint16 {
	borderNum := (p.orientation.rotations + 1) % 4
	isBorderFlipped := borderNum > 1
	if p.orientation.flipped {
		borderNum = (7 - p.orientation.rotations) % 4
		isBorderFlipped = borderNum < 2
	}

	if isBorderFlipped {
		return p.tile.borders[1][borderNum]
	} else {
		return p.tile.borders[0][borderNum]
	}
}

func (p *photoTilePlacement) getBottomBorder() uint16 {
	borderNum := (p.orientation.rotations + 2) % 4
	isBorderFlipped := borderNum < 2
	if p.orientation.flipped {
		borderNum = (6 - p.orientation.rotations) % 4
		isBorderFlipped = borderNum > 1
	}

	if isBorderFlipped {
		return p.tile.borders[1][borderNum]
	} else {
		return p.tile.borders[0][borderNum]
	}
}

func (p *photoTilePlacement) getLeftBorder() uint16 {
	borderNum := (p.orientation.rotations + 3) % 4
	isBorderFlipped := borderNum < 2
	if p.orientation.flipped {
		borderNum = (5 - p.orientation.rotations) % 4
		isBorderFlipped = borderNum > 1
	}

	if isBorderFlipped {
		return p.tile.borders[1][borderNum]
	} else {
		return p.tile.borders[0][borderNum]
	}
}

func (p *photoTilePlacement) getValue(row, col int) bool {
	r, c := p.orientation.translate(row, col)
	return p.tile.text[r][c] == '#'
}

type Grid [photoGridSize][photoGridSize]*photoTilePlacement

func (g Grid) GetCornerProduct() int {
	return g[0][0].tile.tileID *
		g[0][photoGridSize-1].tile.tileID *
		g[photoGridSize-1][0].tile.tileID *
		g[photoGridSize-1][photoGridSize-1].tile.tileID
}
