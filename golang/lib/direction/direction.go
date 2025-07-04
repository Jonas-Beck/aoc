package direction

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d Direction) GetOffset() (int, int) {
	switch d {
	case UP:
		return -1, 0
	case RIGHT:
		return 0, 1
	case DOWN:
		return 1, 0
	case LEFT:
		return 0, -1
	default:
		return 0, 0
	}
}

func (d Direction) Next() Direction {
	return Direction((int(d) + 1) % 4)
}
