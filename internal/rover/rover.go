package rover

import "fmt"

type Rover struct {
	grid Grid
}

func (r *Rover) ExecuteCommands(commands string) string {
	facing := North
	ok := true

	for _, command := range commands {
		switch command {
		case 'L':
			facing = facing.rotateLeft()
		case 'R':
			facing = facing.rotateRight()
		case 'M':
			ok = r.grid.MoveForwards(facing)
			if !ok {
				break
			}
		}
	}
	x, y := r.grid.Location()
	var hitObstaclePrefix string
	if !ok {
		hitObstaclePrefix = "O:"
	}

	return fmt.Sprintf("%s%d:%d:%c", hitObstaclePrefix, x, y, facing)
}

func New(grid Grid) *Rover {
	return &Rover{grid: grid}
}
