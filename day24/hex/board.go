package hex

type Board map[Coordinate]bool

func (b Board) Flip(coordinate Coordinate) {
	if b[coordinate] {
		delete(b, coordinate)
	} else {
		b[coordinate] = true
	}
}

func (b Board) CountBlack() int {
	return len(b)
}

func (b *Board) Step() {
	newBoard := make(Board)
	countBlackNeighbors := make(map[Coordinate]int)
	for blackCoord := range *b {
		var countNeighbors int
		for _, dir := range AllDirections {
			c := blackCoord
			c.Move(dir)
			if (*b)[c] {
				countNeighbors++
			} else {
				countBlackNeighbors[c]++
				if countBlackNeighbors[c] == 2 {
					newBoard[c] = true
				} else if countBlackNeighbors[c] > 2 {
					delete(newBoard, c)
				}
			}
		}

		if countNeighbors == 1 || countNeighbors == 2 {
			newBoard[blackCoord] = true
		}
	}

	*b = newBoard
}
