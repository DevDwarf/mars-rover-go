package grid_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid"
	"testing"
)

func TestSizeFrom(t *testing.T) {
	t.Run("succeeds creating a size struct with the correct values", func(t *testing.T) {
		size, err := grid.SizeFrom(100, 100)
		assert.NoError(t, err)

		assert.Equal(t, width, size.Width)
		assert.Equal(t, height, size.Height)
	})

	t.Run("fails when passed a width less than 1", func(t *testing.T) {
		_, err := grid.SizeFrom(0, 100)
		assert.EqualError(t, err, grid.ErrInvalidWidth.Error())
	})

	t.Run("fails when passed a width greater than 500", func(t *testing.T) {
		_, err := grid.SizeFrom(501, 100)
		assert.EqualError(t, err, grid.ErrInvalidWidth.Error())
	})

	t.Run("fails when passed a height less than 1", func(t *testing.T) {
		_, err := grid.SizeFrom(100, 0)
		assert.EqualError(t, err, grid.ErrInvalidHeight.Error())
	})

	t.Run("fails when passed a height greater than 500", func(t *testing.T) {
		_, err := grid.SizeFrom(100, 501)
		assert.EqualError(t, err, grid.ErrInvalidHeight.Error())
	})

	t.Run("fails when passed a width less than 1 and a height less than 1", func(t *testing.T) {
		_, err := grid.SizeFrom(0, 0)
		assert.EqualError(t, err, grid.ErrInvalidWidth.Error())
	})

	t.Run("fails when passed a width greater than 500 and a height greater than 500", func(t *testing.T) {
		_, err := grid.SizeFrom(501, 501)
		assert.EqualError(t, err, grid.ErrInvalidWidth.Error())
	})
}
