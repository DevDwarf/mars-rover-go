package grid_object

type Options struct {
	Position struct {
		X int
		Y int
	}
	Direction   string
	IsMovable   bool
	IsRotatable bool
	Type        Type
}

type GridObject struct {
	IsGridObject bool
	IsMovable    bool
	IsRotatable  bool
	Position     *Position
	Direction    *Direction
	Type         Type
}

func New(options Options) (*GridObject, error) {
	direction, err := DirectionFrom(options.Direction)
	if err != nil {
		return nil, err
	}

	return &GridObject{
		IsGridObject: true,
		IsMovable:    options.IsMovable,
		IsRotatable:  options.IsRotatable,
		Position:     PositionFrom(options.Position.X, options.Position.Y),
		Direction:    direction,
		Type:         options.Type,
	}, nil
}
