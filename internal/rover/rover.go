package rover

import "fmt"

type Rover struct {
}

func (r *Rover) ExecuteCommands(commands string) string {
	x := 0
	y := 0
	facing := North

	for _, command := range commands {
		switch command {
		case 'L':
			facing = facing.rotateLeft()
		case 'R':
			facing = facing.rotateRight()
		case 'M':
			x, y = r.moveForwards(facing, x, y)
		}
	}

	return fmt.Sprintf("%d:%d:%c", x, y, facing)
}

func (r *Rover) moveForwards(facing direction, x, y int) (int, int) {
	gridSize := 10

	switch facing {
	case North:
		y = (y + 1) % gridSize
	case South:
		y = (y - 1 + gridSize) % gridSize
	case East:
		x = (x + 1) % gridSize
	case West:
		x = (x - 1 + gridSize) % gridSize
	}
	return x, y
}

func New(_ Grid) *Rover {
	return &Rover{}
}
