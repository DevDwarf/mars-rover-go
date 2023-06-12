package grid_object

import "errors"

var (
	ErrGridObjectNotRotatable = errors.New("grid object is not rotatable")
)

func (g *GridObject) RotateLeft() error {
	if !g.IsRotatable {
		return ErrGridObjectNotRotatable
	}

	g.Direction.RotateLeft()

	return nil
}

func (g *GridObject) RotateRight() error {
	if !g.IsRotatable {
		return ErrGridObjectNotRotatable
	}

	g.Direction.RotateRight()

	return nil
}
