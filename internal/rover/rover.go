package rover

import "fmt"

type Rover struct {
	grid   Grid
	facing direction
}

func (r *Rover) ExecuteCommands(commands string) string {
	moveResult := r.executeCommands(commands)

	return r.displayResult(moveResult)
}

func (r *Rover) executeCommands(commands string) MoveResult {
	for _, command := range commands {
		switch command {
		case 'L':
			r.facing = r.facing.rotateLeft()
		case 'R':
			r.facing = r.facing.rotateRight()
		case 'M':
			result := r.grid.MoveForwards(r.facing)
			if result != MoveOk {
				return result
			}
		}
	}
	return MoveOk
}

func (r *Rover) displayResult(moveResult MoveResult) string {
	location := r.grid.Location()

	var hitObstaclePrefix string
	if moveResult == HitObstacle {
		hitObstaclePrefix = "O:"
	}

	return fmt.Sprintf("%s%d:%d:%c",
		hitObstaclePrefix,
		location.x,
		location.y,
		r.facing,
	)
}

func New(grid Grid) *Rover {
	return &Rover{
		grid:   grid,
		facing: North,
	}
}
