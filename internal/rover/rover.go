package rover

import "fmt"

type Rover struct {
}

func (r *Rover) ExecuteCommands(commands string) string {
	y := 0
	direction := North

	for _, command := range commands {
		switch command {
		case 'L':
			direction = direction.rotateLeft()
		case 'R':
			direction = direction.rotateRight()
		}
	}

	return fmt.Sprintf("0:%d:%c", y, direction)
}

func New() *Rover {
	return &Rover{}
}
