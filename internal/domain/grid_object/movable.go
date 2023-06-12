package grid_object

import "errors"

var (
	ErrGridObjectNotMovable = errors.New("grid object is not movable")
)

func (g *GridObject) MoveNorth() error {
	if !g.IsMovable {
		return ErrGridObjectNotMovable
	}

	g.Position.IncrementY()

	return nil
}

func (g *GridObject) MoveSouth() error {
	if !g.IsMovable {
		return ErrGridObjectNotMovable
	}

	g.Position.DecrementY()

	return nil
}

func (g *GridObject) MoveEast() error {
	if !g.IsMovable {
		return ErrGridObjectNotMovable
	}

	g.Position.IncrementX()

	return nil
}

func (g *GridObject) MoveWest() error {
	if !g.IsMovable {
		return ErrGridObjectNotMovable
	}

	g.Position.DecrementX()

	return nil
}
