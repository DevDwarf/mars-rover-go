package grid_object

import "errors"

var (
	ErrInvalidDirection = errors.New("direction is invalid")
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var Directions = map[Direction]string{
	North: "N",
	East:  "E",
	South: "S",
	West:  "W",
}

func DirectionFrom(direction string) (*Direction, error) {
	for k, v := range Directions {
		if v == direction {
			return &k, nil
		}
	}

	return nil, ErrInvalidDirection
}

// RotateRight rotates the direction to the right
// We increment the direction by 1 and then modulo 4 to ensure we don't go out of bounds
func (d *Direction) RotateRight() {
	// This method only works because the directions are in order
	// N -> E -> S -> W -> N
	// 0 -> 1 -> 2 -> 3 -> 0
	*d = (*d + 1) % 4
}

// RotateLeft rotates the direction to the left
// We decrement the direction by 1 then add 4 and then modulo 4 to ensure we don't go out of bounds
// The add 4 is to ensure we don't get a negative number
func (d *Direction) RotateLeft() {
	// This method only works because the directions are in order
	*d = (*d - 1 + 4) % 4
}

// add String function to Direction
func (d *Direction) String() string {
	return Directions[*d]
}
