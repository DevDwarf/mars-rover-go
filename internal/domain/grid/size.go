package grid

import (
	"errors"
	"strconv"
)

var (
	MaxWidth         = 500
	MaxHeight        = 500
	ErrInvalidWidth  = errors.New("the width must be greater than 0 and less than or equal to " + strconv.Itoa(MaxWidth))
	ErrInvalidHeight = errors.New("the height must be greater than 0 and less than or equal to " + strconv.Itoa(MaxHeight))
)

type Size struct {
	Width  int
	Height int
}

func SizeFrom(width, height int) (*Size, error) {
	if width <= 0 || width > MaxWidth {
		return nil, ErrInvalidWidth
	}

	if height <= 0 || height > MaxHeight {
		return nil, ErrInvalidHeight
	}

	return &Size{
		Width:  width,
		Height: height,
	}, nil
}
