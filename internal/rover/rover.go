package rover

import "fmt"

type Rover struct {
	grid *Grid
}

func (r *Rover) ExecuteCommands(commands string) string {
	facing := North

	for _, command := range commands {
		switch command {
		case 'L':
			facing = facing.rotateLeft()
		case 'R':
			facing = facing.rotateRight()
		case 'M':
			r.grid.MoveForwards(facing)
		}
	}
	x, y := r.grid.Location()

	return fmt.Sprintf("%d:%d:%c", x, y, facing)
}

func New(grid *Grid) *Rover {
	return &Rover{grid: grid}
}
