package rover

type Grid interface {
	// Location is the final Rover Location.
	Location() Location

	// MoveForwards moves Rover in the given direction and indicate what occurred.
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
		obstacles: make(map[Location]bool),
	}
}

type SquareGrid struct {
	location  Location
	size      int
	obstacles map[Location]bool
}

type Location struct {
	x int
	y int
}

func (g *SquareGrid) Location() Location {
	return Location{x: g.location.x, y: g.location.y}
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

func (g *SquareGrid) requestedMovement(facing direction) Location {
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

	return Location{x: x, y: y}
}

func (g *SquareGrid) AddObstacleAt(x int, y int) {
	g.obstacles[Location{x, y}] = true
}
