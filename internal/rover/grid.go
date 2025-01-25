package rover

type Grid interface {
	// Location is the final Rover location.
	Location() (x int, y int)

	// MoveForwards moves Rover in the given direction and returns true if succeeds.
	MoveForwards(facing direction) MoveResult
}

type MoveResult bool

const (
	MoveOk      MoveResult = false
	HitObstacle MoveResult = true
)

func NewSquareGrid() *SquareGrid {
	return &SquareGrid{
		size:      10,
		obstacles: make(map[location]bool),
	}
}

type SquareGrid struct {
	x         int
	y         int
	size      int
	obstacles map[location]bool
}

type location struct {
	x int
	y int
}

func (g *SquareGrid) Location() (x int, y int) {
	x = g.x
	y = g.y
	return
}

func (g *SquareGrid) MoveForwards(facing direction) MoveResult {
	x := g.x
	y := g.y
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

	hitObstacle := g.obstacles[location{x, y}]
	if hitObstacle {
		return HitObstacle
	}

	g.x = x
	g.y = y
	return MoveOk
}

func (g *SquareGrid) AddObstacleAt(x int, y int) {
	g.obstacles[location{x, y}] = true
}
