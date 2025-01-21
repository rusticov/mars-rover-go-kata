package rover

import "fmt"

type Direction rune

const (
	North Direction = 'N'
	East  Direction = 'E'
	South Direction = 'S'
	West  Direction = 'W'
)

func (d Direction) rotateRight() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	default:
		return North
	}
}

type Rover struct {
}

func (r *Rover) ExecuteCommands(commands string) string {
	direction := North

	for _, command := range commands {
		switch command {
		case 'R':
			direction = direction.rotateRight()
		}
	}

	return fmt.Sprintf("0:0:%c", direction)
}

func New() *Rover {
	return &Rover{}
}
