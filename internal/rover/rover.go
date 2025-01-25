package rover

import "fmt"

type Rover struct {
	grid   Grid
	facing direction
}

func (r *Rover) ExecuteCommands(commands string) string {
	var moveResult MoveResult

	for _, command := range commands {
		switch command {
		case 'L':
			r.facing = r.facing.rotateLeft()
		case 'R':
			r.facing = r.facing.rotateRight()
		case 'M':
			moveResult = r.grid.MoveForwards(r.facing)
			if moveResult != MoveOk {
				break
			}
		}
	}

	return r.displayResult(r.facing, moveResult)
}

func (r *Rover) displayResult(facing direction, moveResult MoveResult) string {
	x, y := r.grid.Location()

	var hitObstaclePrefix string
	if moveResult == HitObstacle {
		hitObstaclePrefix = "O:"
	}

	return fmt.Sprintf("%s%d:%d:%c", hitObstaclePrefix, x, y, facing)
}

func New(grid Grid) *Rover {
	return &Rover{
		grid:   grid,
		facing: North,
	}
}
