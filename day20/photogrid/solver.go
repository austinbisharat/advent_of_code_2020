package photogrid

import "log"

type photoGridSolver struct {
	tiles []PhotoTile
	grid  Grid

	placements map[*photoTilePlacement]bool

	// TODO: Definitly a way to factor this to use just one map and do rotations on the fly
	topBorders  map[uint16]map[*photoTilePlacement]bool
	leftBorders map[uint16]map[*photoTilePlacement]bool
}

func NewPhotoGridSolver(tiles []PhotoTile) *photoGridSolver {
	if len(tiles) != photoGridSize*photoGridSize {
		log.Fatal("Incorrect number of tiles")
	}
	solver := photoGridSolver{
		tiles: tiles,
	}
	solver.init()
	return &solver
}

func (solver *photoGridSolver) init() {
	solver.topBorders = make(map[uint16]map[*photoTilePlacement]bool)
	solver.leftBorders = make(map[uint16]map[*photoTilePlacement]bool)
	solver.placements = make(map[*photoTilePlacement]bool)

	for i := range solver.tiles {
		for rotation := uint8(0); rotation < 4; rotation++ {
			for flipped := 0; flipped < 2; flipped++ {
				placement := photoTilePlacement{
					tile: &solver.tiles[i],
					orientation: orientation{
						flipped:   flipped%2 == 0,
						rotations: rotation,
						size:      tileSize,
					},
				}

				topBorder := placement.getTopBorder()
				if solver.topBorders[topBorder] == nil {
					solver.topBorders[topBorder] = make(map[*photoTilePlacement]bool)
				}
				solver.topBorders[topBorder][&placement] = true

				leftBorder := placement.getLeftBorder()
				if solver.leftBorders[leftBorder] == nil {
					solver.leftBorders[leftBorder] = make(map[*photoTilePlacement]bool)
				}
				solver.leftBorders[leftBorder][&placement] = true

				solver.placements[&placement] = true
			}
		}
	}
}

func (solver *photoGridSolver) Solve() []Grid {
	var solutions []Grid
	solver.solve(make(map[*PhotoTile]bool), &solutions)
	return solutions
}

func (solver *photoGridSolver) solve(tilesPlaced map[*PhotoTile]bool, solutions *[]Grid) {
	nextTileNum := len(tilesPlaced)
	if nextTileNum == photoGridSize*photoGridSize {
		*solutions = append(*solutions, solver.grid)
		return
	}

	// special case for when there are no restrictions on first tile
	if nextTileNum == 0 {
		for placement := range solver.placements {
			solver.grid[0][0] = placement
			tilesPlaced[placement.tile] = true

			solver.solve(tilesPlaced, solutions)

			delete(tilesPlaced, placement.tile)
			solver.grid[0][0] = nil
		}
		return
	}

	row := nextTileNum / photoGridSize
	col := nextTileNum % photoGridSize

	if row == 0 {
		prevPlacement := solver.grid[row][col-1]
		for placement := range solver.leftBorders[prevPlacement.getRightBorder()] {
			if tilesPlaced[placement.tile] {
				continue
			}
			solver.grid[row][col] = placement
			tilesPlaced[placement.tile] = true

			solver.solve(tilesPlaced, solutions)

			delete(tilesPlaced, placement.tile)
			solver.grid[row][col] = nil
		}
		return
	}

	if col == 0 {
		prevPlacement := solver.grid[row-1][col]
		for placement := range solver.topBorders[prevPlacement.getBottomBorder()] {
			if tilesPlaced[placement.tile] {
				continue
			}
			solver.grid[row][col] = placement
			tilesPlaced[placement.tile] = true

			solver.solve(tilesPlaced, solutions)

			delete(tilesPlaced, placement.tile)
			solver.grid[row][col] = nil
		}
		return
	}

	leftPlacement := solver.grid[row][col-1]
	abovePlacement := solver.grid[row-1][col]

	candidatesLeft := solver.leftBorders[leftPlacement.getRightBorder()]
	candidatesAbove := solver.topBorders[abovePlacement.getBottomBorder()]

	for placement := range candidatesLeft {
		if tilesPlaced[placement.tile] || !candidatesAbove[placement] {
			continue
		}

		solver.grid[row][col] = placement
		tilesPlaced[placement.tile] = true

		solver.solve(tilesPlaced, solutions)

		delete(tilesPlaced, placement.tile)
		solver.grid[row][col] = nil
	}

	return
}
