package grid_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid"
	"mars-rover-go/internal/domain/grid_object"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("succeeds creating a new grid with correct size", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		assert.Equal(t, width, g.Size.Width)
		assert.Equal(t, height, g.Size.Height)
	})

	t.Run("returns an error when SizeFrom returns an error", func(t *testing.T) {
		_, err := grid.New(0, height)
		assert.Error(t, err)
	})
}

func TestIsValidPosition(t *testing.T) {
	t.Run("returns no errors on correct position", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		errs := g.IsValidPosition(validX, validY)
		assert.Empty(t, errs)
	})

	t.Run("returns expected error when x is invalid", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		e1 := g.IsValidPosition(-1, validY)
		assert.Equal(t, []error{grid.ErrOutOfBoundsX}, e1)

		e3 := g.IsValidPosition(width+1, validY)
		assert.Equal(t, []error{grid.ErrOutOfBoundsX}, e3)
	})

	t.Run("returns expected error when y is invalid", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		e1 := g.IsValidPosition(validX, -1)
		assert.Equal(t, []error{grid.ErrOutOfBoundsY}, e1)

		e3 := g.IsValidPosition(validX, height+1)
		assert.Equal(t, []error{grid.ErrOutOfBoundsY}, e3)
	})

	t.Run("returns expected errors when x and y are invalid", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		e1 := g.IsValidPosition(-1, -1)
		assert.Equal(t, []error{grid.ErrOutOfBoundsX, grid.ErrOutOfBoundsY}, e1)

		e3 := g.IsValidPosition(width+1, height+1)
		assert.Equal(t, []error{grid.ErrOutOfBoundsX, grid.ErrOutOfBoundsY}, e3)
	})
}

func TestFindObjectByCoordinates(t *testing.T) {
	t.Run("returns nil when no object is found", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		assert.Nil(t, g.FindObjectByCoordinates(validX, validY))
	})

	t.Run("returns the object when found", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, err := grid_object.New(grid_object.Options{
			Position:  objectOptions1.Position,
			Direction: objectOptions1.Direction,
			Type:      objectOptions1.Type,
		})
		assert.NoError(t, err)

		g.Objects = append(g.Objects, o)

		assert.Equal(t, o, g.FindObjectByCoordinates(objectOptions1.Position.X, objectOptions1.Position.Y))
	})
}

func TestListObjectsByType(t *testing.T) {
	t.Run("returns an empty slice when no objects are found", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		assert.Empty(t, g.ListObjectsByType(grid_object.Rover))
	})

	t.Run("returns the correct objects when found", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o1, err := grid_object.New(grid_object.Options{
			Position:  objectOptions1.Position,
			Direction: objectOptions1.Direction,
			Type:      objectOptions1.Type,
		})
		assert.NoError(t, err)

		o2, err := grid_object.New(grid_object.Options{
			Position:  objectoOptions2.Position,
			Direction: objectoOptions2.Direction,
			Type:      objectoOptions2.Type,
		})
		assert.NoError(t, err)

		o3, err := grid_object.New(grid_object.Options{
			Position:  objectOptions3.Position,
			Direction: objectOptions3.Direction,
			Type:      objectOptions3.Type,
		})
		assert.NoError(t, err)

		g.Objects = append(g.Objects, o1, o2, o3)

		assert.Equal(t, []*grid_object.GridObject{o1, o2}, g.ListObjectsByType(grid_object.Rover))
	})
}

func TestAddObject(t *testing.T) {
	t.Run("succeeds adding a new object", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddObject(grid_object.Options{
			Position:  objectOptions1.Position,
			Direction: objectOptions1.Direction,
			Type:      objectOptions1.Type,
		})
		assert.Empty(t, errs)

		assert.Equal(t, []*grid_object.GridObject{o}, g.Objects)
	})

	t.Run("returns an error when object is out of bounds", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		_, errs := g.AddObject(grid_object.Options{
			Position:  grid_object.Position{X: width + 1, Y: height + 1},
			Direction: objectOptions1.Direction,
			Type:      objectOptions1.Type,
		})
		assert.Equal(t, []error{grid.ErrOutOfBoundsX, grid.ErrOutOfBoundsY}, errs)
	})

	t.Run("returns an error when direction is invalid", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		_, errs := g.AddObject(grid_object.Options{
			Position:  objectOptions1.Position,
			Direction: "invalid",
			Type:      objectOptions1.Type,
		})
		assert.Equal(t, []error{grid_object.ErrInvalidDirection}, errs)
	})
}

func TestAddRover(t *testing.T) {
	t.Run("succeeds adding a new rover", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, objectOptions1.Direction)
		assert.Empty(t, errs)

		assert.Equal(t, []*grid_object.GridObject{o}, g.Objects)
	})

	t.Run("returns an error when AddObject fails", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		_, errs := g.AddRover(width+1, height+1, objectOptions1.Direction)
		assert.Equal(t, []error{grid.ErrOutOfBoundsX, grid.ErrOutOfBoundsY}, errs)
	})
}

func TestAddBoulder(t *testing.T) {
	t.Run("succeeds adding a new boulder", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, objectOptions1.Direction)
		assert.Empty(t, errs)

		assert.Equal(t, []*grid_object.GridObject{o}, g.Objects)
	})

	t.Run("returns an error when AddObject fails", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		_, errs := g.AddRover(width+1, height+1, objectOptions1.Direction)
		assert.Equal(t, []error{grid.ErrOutOfBoundsX, grid.ErrOutOfBoundsY}, errs)
	})
}
