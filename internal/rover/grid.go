package rover

type Grid interface {
	Location() (x int, y int)
	MoveForwards(facing direction)
}

type SquareGrid struct {
	x    int
	y    int
	size int
}

func (g *SquareGrid) Location() (x int, y int) {
	x = g.x
	y = g.y
	return
}

func (g *SquareGrid) MoveForwards(facing direction) {
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
}

func NewSquareGrid() *SquareGrid {
	return &SquareGrid{
		size: 10,
	}
}
