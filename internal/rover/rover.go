package rover

import "fmt"

type Rover struct {
}

func (r *Rover) ExecuteCommands(commands string) string {
	gridSize := 10

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
		}
	}

	return fmt.Sprintf("%d:%d:%c", x, y, facing)
}

func New() *Rover {
	return &Rover{}
}
