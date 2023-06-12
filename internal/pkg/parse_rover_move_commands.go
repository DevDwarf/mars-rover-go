package pkg

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrorInvalidRoverMoveParameters = errors.New("invalid rover move parameters")
)

type Actions []string

func ParseRoverMoveCommands(str string) ([]string, error) {
	// Define the regular expression pattern
	pattern := `^[LRM]+$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Validate the string against the regular expression
	if regex.MatchString(str) {
		// Create the Actions struct
		actions := strings.Split(str, "")

		return actions, nil
	}

	return nil, ErrorInvalidRoverMoveParameters
}
