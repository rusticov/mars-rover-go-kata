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
	location  location
	size      int
	obstacles map[location]bool
}

type location struct {
	x int
	y int
}

func (g *SquareGrid) Location() (x int, y int) {
	x = g.location.x
	y = g.location.y
	return
}

func (g *SquareGrid) MoveForwards(facing direction) MoveResult {
	endLocation := g.requestedMovement(facing)

	hitObstacle := g.obstacles[endLocation]
	if hitObstacle {
		return HitObstacle
	}

	g.location = endLocation
	return MoveOk
}

func (g *SquareGrid) requestedMovement(facing direction) location {
	x := g.location.x
	y := g.location.y
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

	return location{x: x, y: y}
}

func (g *SquareGrid) AddObstacleAt(x int, y int) {
	g.obstacles[location{x, y}] = true
}
