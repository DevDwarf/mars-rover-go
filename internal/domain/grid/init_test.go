package grid_test

import (
	"mars-rover-go/internal/domain/grid_object"
)

type ObjectTestOptions = struct {
	Position  struct{ X, Y int }
	Direction string
	Type      grid_object.Type
}

var (
	width          = 100
	height         = 100
	validX         = 50
	validY         = 50
	objectOptions1 = ObjectTestOptions{
		Position:  struct{ X, Y int }{X: 10, Y: 10},
		Direction: "N",
		Type:      grid_object.Rover,
	}
	objectoOptions2 = ObjectTestOptions{
		Position:  struct{ X, Y int }{X: 64, Y: 35},
		Direction: "S",
		Type:      grid_object.Rover,
	}
	objectOptions3 = ObjectTestOptions{
		Position:  struct{ X, Y int }{X: 54, Y: 23},
		Direction: "N",
		Type:      grid_object.Boulder,
	}
)
