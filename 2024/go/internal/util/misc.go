package util

// IntAbs returns the absolute value of an integer
func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Coordinate struct {
	X, Y int
}

// Movement describes NESW movement
type Movement int

const (
	MoveUp Movement = iota
	MoveRight
	MoveDown
	MoveLeft
	MoveUnknown
)

func (m Movement) CycleClockwise() Movement {
	switch m {
	case MoveUp:
		return MoveRight
	case MoveRight:
		return MoveDown
	case MoveDown:
		return MoveLeft
	case MoveLeft:
		return MoveUp
	default:
		return MoveUnknown
	}
}

func (m Movement) String() string {
	switch m {
	case MoveUp:
		return "up"
	case MoveDown:
		return "down"
	case MoveLeft:
		return "left"
	case MoveRight:
		return "right"
	default:
		return "unknown"
	}
}
