package photogrid

type orientation struct {
	flipped   bool
	rotations uint8
	size      int
}

func (o *orientation) translate(x, y int) (int, int) {
	x, y = 2*x-o.size+1, 2*y-o.size+1
	for i := uint8(0); i < o.rotations; i++ {
		x, y = y, -1*x
	}
	if o.flipped {
		y *= -1
	}
	return (x + o.size - 1) / 2, (y + o.size - 1) / 2
}
