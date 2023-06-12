package grid_object_test

import (
	"github.com/stretchr/testify/assert"
	"mars-rover-go/internal/domain/grid_object"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should return a new grid object", func(t *testing.T) {
		_, err := grid_object.New(grid_object.Options{
			Position:  objectoOptions2.Position,
			Direction: objectoOptions2.Direction,
			Type:      objectoOptions2.Type,
		})
		assert.NoError(t, err)
	})

	t.Run("should return an error if the direction is invalid", func(t *testing.T) {
		_, err := grid_object.New(grid_object.Options{
			Position:  objectoOptions2.Position,
			Direction: "Z",
			Type:      objectoOptions2.Type,
		})
		assert.EqualError(t, err, grid_object.ErrInvalidDirection.Error())
	})
}
