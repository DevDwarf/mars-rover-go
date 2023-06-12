package grid_object

type Position struct {
	X int
	Y int
}

// PositionFrom creates a new Position from the given x and y coordinates
// No validation because it gets handled on the Grid level
// We can assume that the coordinates are valid at this point
func PositionFrom(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) IncrementX() {
	p.X++
}

func (p *Position) DecrementX() {
	p.X--
}

func (p *Position) IncrementY() {
	p.Y++
}

func (p *Position) DecrementY() {
	p.Y--
}
