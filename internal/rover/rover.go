package rover

import "fmt"

type Rover struct {
}

func (r *Rover) ExecuteCommands(commands string) string {
	direction := North

	for _, command := range commands {
		switch command {
		case 'L':
			direction = direction.rotateLeft()
		case 'R':
			direction = direction.rotateRight()
		}
	}

	return fmt.Sprintf("0:0:%c", direction)
}

func New() *Rover {
	return &Rover{}
}
