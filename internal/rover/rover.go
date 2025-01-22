package rover

import "fmt"

type Rover struct {
}

func (r *Rover) ExecuteCommands(commands string) string {
	gridSize := 10

	y := 0
	facing := North

	for _, command := range commands {
		switch command {
		case 'L':
			facing = facing.rotateLeft()
		case 'R':
			facing = facing.rotateRight()
		case 'M':
			y = (y + 1) % gridSize
		}
	}

	return fmt.Sprintf("0:%d:%c", y, facing)
}

func New() *Rover {
	return &Rover{}
}
