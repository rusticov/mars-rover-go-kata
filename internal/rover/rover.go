package rover

import "fmt"

type Rover struct {
	grid Grid
}

func (r *Rover) ExecuteCommands(commands string) string {
	facing := North
	var moveResult MoveResult

	for _, command := range commands {
		switch command {
		case 'L':
			facing = facing.rotateLeft()
		case 'R':
			facing = facing.rotateRight()
		case 'M':
			moveResult = r.grid.MoveForwards(facing)
			if moveResult != MoveOk {
				break
			}
		}
	}

	return r.displayResult(facing, moveResult)
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
	return &Rover{grid: grid}
}
