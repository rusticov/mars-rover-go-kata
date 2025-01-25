package rover

type Grid struct {
	x    int
	y    int
	size int
}

func (g *Grid) moveForwards(facing direction, x, y int) (int, int) {
	switch facing {
	case North:
		y = (y + 1) % g.size
	case South:
		y = (y - 1 + g.size) % g.size
	case East:
		x = (x + 1) % g.size
	case West:
		x = (x - 1 + g.size) % g.size
	}
	return x, y
}

func NewGrid() *Grid {
	return &Grid{
		size: 10,
	}
}
