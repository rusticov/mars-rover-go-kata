package rover

type direction rune

const (
	North direction = 'N'
	East  direction = 'E'
	South direction = 'S'
	West  direction = 'W'
)

func (d direction) rotateRight() direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	default:
		return North
	}
}

func (d direction) rotateLeft() direction {
	switch d {
	case North:
		return West
	case West:
		return South
	case South:
		return East
	default:
		return North
	}
}
