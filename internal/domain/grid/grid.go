package grid

import (
	"errors"
	"mars-rover-go/internal/domain/grid_object"
)

var (
	ErrOutOfBoundsX = errors.New("x coordinate is out of bounds")
	ErrOutOfBoundsY = errors.New("y coordinate is out of bounds")
)

type Grid struct {
	Size    *Size
	Objects []*grid_object.GridObject
}

func New(width, height int) (*Grid, error) {
	s, err := SizeFrom(width, height)
	if err != nil {
		return nil, err
	}

	return &Grid{
		Size: s,
	}, nil
}

// IsValidX checks if the given x coordinate is within the grid's width
// Unexported because it's only used internally
func (g *Grid) isValidX(x int) error {
	if x < 0 || x > g.Size.Width {
		return ErrOutOfBoundsX
	}

	return nil
}

// IsValidY checks if the given y coordinate is within the grid's height
// Unexported because it's only used internally
func (g *Grid) isValidY(y int) error {
	if y < 0 || y > g.Size.Height {
		return ErrOutOfBoundsY
	}

	return nil
}

// IsValidPosition checks if the given x and y coordinates are within the grid's width and height
// Exported because it's used by other packages
func (g *Grid) IsValidPosition(x, y int) []error {
	var errs []error

	if err := g.isValidX(x); err != nil {
		errs = append(errs, err)
	}

	if err := g.isValidY(y); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func (g *Grid) FindObjectByCoordinates(x, y int) *grid_object.GridObject {
	for _, o := range g.Objects {
		if o.Position.X == x && o.Position.Y == y {
			return o
		}
	}

	return nil
}

func (g *Grid) ListObjectsByType(t grid_object.Type) []*grid_object.GridObject {
	objects := make([]*grid_object.GridObject, 0)

	for _, o := range g.Objects {
		if o.Type == t {
			objects = append(objects, o)
		}
	}

	return objects
}

func (g *Grid) AddObject(options grid_object.Options) (*grid_object.GridObject, []error) {
	if errs := g.IsValidPosition(options.Position.X, options.Position.Y); len(errs) > 0 {
		return nil, errs
	}

	o, err := grid_object.New(options)
	if err != nil {
		return nil, []error{err}
	}

	g.Objects = append(g.Objects, o)

	return o, nil
}

func (g *Grid) AddRover(x, y int, direction string) (*grid_object.GridObject, []error) {
	o, errs := g.AddObject(grid_object.Options{
		Position: struct {
			X int
			Y int
		}{
			X: x,
			Y: y,
		},
		Direction:   direction,
		IsRotatable: true,
		IsMovable:   true,
		Type:        grid_object.Rover,
	})
	if len(errs) > 0 {
		return nil, errs
	}

	return o, nil
}

func (g *Grid) AddBoulder(x, y int, direction string) (*grid_object.GridObject, []error) {
	o, errs := g.AddObject(grid_object.Options{
		Position: struct {
			X int
			Y int
		}{
			X: x,
			Y: y,
		},
		Direction: direction,
		Type:      grid_object.Boulder,
	})
	if len(errs) > 0 {
		return nil, errs
	}

	return o, nil
}
