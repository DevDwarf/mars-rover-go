package grid_object_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid"
	"mars-rover-go/internal/domain/grid_object"
	"testing"
)

func TestMoveNorth(t *testing.T) {
	t.Run("should move north", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveNorth()
		assert.NoError(t, err)

		assert.Equal(t, objectOptions1.Position.Y+1, o.Position.Y)
	})

	t.Run("returns expected error when trying to move", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveNorth()
		assert.EqualError(t, err, grid_object.ErrGridObjectNotMovable.Error())
	})
}

func TestMoveSouth(t *testing.T) {
	t.Run("should move south", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveSouth()
		assert.NoError(t, err)

		assert.Equal(t, objectOptions1.Position.Y-1, o.Position.Y)
	})

	t.Run("returns expected error when trying to move", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveSouth()
		assert.EqualError(t, err, grid_object.ErrGridObjectNotMovable.Error())
	})
}

func TestMoveEast(t *testing.T) {
	t.Run("should move east", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveEast()
		assert.NoError(t, err)

		assert.Equal(t, objectOptions1.Position.X+1, o.Position.X)
	})

	t.Run("returns expected error when trying to move", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveEast()
		assert.EqualError(t, err, grid_object.ErrGridObjectNotMovable.Error())
	})
}

func TestMoveWest(t *testing.T) {
	t.Run("should move west", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveWest()
		assert.NoError(t, err)

		assert.Equal(t, objectOptions1.Position.X-1, o.Position.X)
	})

	t.Run("returns expected error when trying to move", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.MoveWest()
		assert.EqualError(t, err, grid_object.ErrGridObjectNotMovable.Error())
	})
}
