package hex

type Direction string

const (
	DirNorthEast Direction = "ne"
	DirNorthWest Direction = "nw"
	DirWest      Direction = "w"
	DirEast      Direction = "e"
	DirSouthEast Direction = "se"
	DirSouthWest Direction = "sw"
)

var AllDirections = []Direction{DirNorthEast, DirNorthWest, DirWest, DirEast, DirSouthEast, DirSouthWest}

type Coordinate struct {
	x int
	y int
}

//        (0,0),    (1, 0)
//    (-1, 1), (0, 1)
//        (0, 2)  (1, 2)
func (c *Coordinate) Move(dir Direction) {
	switch dir {
	case DirNorthEast:
		c.x += mod(c.y, 2)
		c.y--
	case DirNorthWest:
		c.x -= mod(c.y+1, 2)
		c.y--
	case DirEast:
		c.x++
	case DirWest:
		c.x--
	case DirSouthEast:
		c.x += mod(c.y, 2)
		c.y++
	case DirSouthWest:
		c.x -= mod(c.y+1, 2)
		c.y++
	default:
		panic("unknown dir")

	}
}

func (c *Coordinate) FollowPath(path Path) {
	for _, dir := range path {
		c.Move(dir)
	}
}

func mod(a, b int) int {
	n := a % b
	if n < 0 {
		n += b
	}
	return n
}

type Path []Direction

func NewPathFromStr(str string) Path {
	var path Path
	for len(str) > 0 {
		dir := string(str[0])
		str = str[1:]
		if dir == "n" || dir == "s" {
			dir += string(str[0])
			str = str[1:]
		}
		path = append(path, Direction(dir))
	}
	return path
}
