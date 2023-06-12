package pkg

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrorInvalidRoverParameters = errors.New("invalid rover parameters")
)

type AddRoverInput struct {
	X   int
	Y   int
	Dir string
}

func ParseAddRoverCommands(size string) (*AddRoverInput, error) {
	// Define the regular expression pattern
	pattern := `^(\d+)\s(\d+)\s([NESW])$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Match the string against the regular expression
	matches := regex.FindStringSubmatch(size)

	if len(matches) == 4 {
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, ErrorInvalidRoverParameters
		}

		y, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, ErrorInvalidRoverParameters
		}

		return &AddRoverInput{
			X:   x,
			Y:   y,
			Dir: matches[3],
		}, nil
	}

	return nil, ErrorInvalidRoverParameters
}
