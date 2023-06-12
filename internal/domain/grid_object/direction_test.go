package grid_object_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid_object"
	"testing"
)

func TestDirectionFrom(t *testing.T) {
	t.Run("should return the correct direction", func(t *testing.T) {
		d, err := grid_object.DirectionFrom("N")
		assert.NoError(t, err)

		assert.Equal(t, grid_object.North, *d)
	})

	t.Run("should return error if direction is invalid", func(t *testing.T) {
		_, err := grid_object.DirectionFrom("Z")
		assert.EqualError(t, err, grid_object.ErrInvalidDirection.Error())
	})
}

func TestRotateRightDirection(t *testing.T) {
	t.Run("should rotate direction right", func(t *testing.T) {
		d := grid_object.North
		d.RotateRight()

		assert.Equal(t, grid_object.East, d)
	})

	t.Run("should rotate direction right and wrap around", func(t *testing.T) {
		d := grid_object.West
		d.RotateRight()

		assert.Equal(t, grid_object.North, d)
	})
}

func TestRotateLeftDirection(t *testing.T) {
	t.Run("should rotate direction left", func(t *testing.T) {
		d := grid_object.North
		d.RotateLeft()

		assert.Equal(t, grid_object.West, d)
	})

	t.Run("should rotate direction left and wrap around", func(t *testing.T) {
		d := grid_object.East
		d.RotateLeft()

		assert.Equal(t, grid_object.North, d)
	})
}
