package rover

type Grid struct {
	x    int
	y    int
	size int
}

func (g *Grid) Location() (x int, y int) {
	x = g.x
	y = g.y
	return
}

func (g *Grid) MoveForwards(facing direction) (int, int) {
	switch facing {
	case North:
		g.y = (g.y + 1) % g.size
	case South:
		g.y = (g.y - 1 + g.size) % g.size
	case East:
		g.x = (g.x + 1) % g.size
	case West:
		g.x = (g.x - 1 + g.size) % g.size
	}
	return g.x, g.y
}

func NewGrid() *Grid {
	return &Grid{
		size: 10,
	}
}
