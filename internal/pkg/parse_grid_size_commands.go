package pkg

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrorInvalidGridSize = errors.New("invalid grid parameters")
)

type GridSizeInput struct {
	X int
	Y int
}

func ParseGridSizeParams(size string) (*GridSizeInput, error) {
	// Define the regular expression pattern
	pattern := `^(\d+)\s(\d+)\s*$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Match the string against the regular expression
	matches := regex.FindStringSubmatch(size)

	if len(matches) == 3 {
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, ErrorInvalidGridSize
		}

		y, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, ErrorInvalidGridSize
		}

		return &GridSizeInput{
			X: x,
			Y: y,
		}, nil
	}

	return nil, ErrorInvalidGridSize
}
