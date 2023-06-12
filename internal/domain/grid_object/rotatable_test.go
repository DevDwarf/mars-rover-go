package grid_object_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid"
	"mars-rover-go/internal/domain/grid_object"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	t.Run("should return error if grid object is not rotatable", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, objectOptions1.Direction)
		assert.Empty(t, errs)

		err = o.RotateLeft()
		assert.EqualError(t, err, grid_object.ErrGridObjectNotRotatable.Error())
	})

	t.Run("should rotate grid object left", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.RotateLeft()
		assert.NoError(t, err)

		assert.Equal(t, o.Direction.String(), "W")
	})
}

func TestRotateRight(t *testing.T) {
	t.Run("should return error if grid object is not rotatable", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddBoulder(objectOptions1.Position.X, objectOptions1.Position.Y, objectOptions1.Direction)
		assert.Empty(t, errs)

		err = o.RotateRight()
		assert.EqualError(t, err, grid_object.ErrGridObjectNotRotatable.Error())
	})

	t.Run("should rotate grid object right", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		err = o.RotateRight()
		assert.NoError(t, err)

		assert.Equal(t, o.Direction.String(), "E")
	})
}
