package grid_object_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid"
	"testing"
)

func TestIncrementX(t *testing.T) {
	t.Run("should increment the x coordinate by 1", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		o.Position.IncrementX()
		assert.Equal(t, o.Position.X, objectOptions1.Position.X+1)
	})
}

func TestDecrementX(t *testing.T) {
	t.Run("should decrement the x coordinate by 1", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		o.Position.DecrementX()
		assert.Equal(t, o.Position.X, objectOptions1.Position.X-1)
	})
}

func TestIncrementY(t *testing.T) {
	t.Run("should increment the y coordinate by 1", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		o.Position.IncrementY()
		assert.Equal(t, o.Position.Y, objectOptions1.Position.Y+1)
	})
}

func TestDecrementY(t *testing.T) {
	t.Run("should decrement the y coordinate by 1", func(t *testing.T) {
		g, err := grid.New(width, height)
		assert.NoError(t, err)

		o, errs := g.AddRover(objectOptions1.Position.X, objectOptions1.Position.Y, "N")
		assert.Empty(t, errs)

		o.Position.DecrementY()
		assert.Equal(t, o.Position.Y, objectOptions1.Position.Y-1)
	})
}
