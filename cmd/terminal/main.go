package main

import (
	"fmt"
	"mars-rover-go/internal/domain/grid"
	"mars-rover-go/internal/domain/grid_object"
	"mars-rover-go/internal/interface/terminal"
	"mars-rover-go/internal/pkg"
)

func main() {
	// Create a new terminal
	term := terminal.New()

	var sessionGrid *grid.Grid

	for {
		fmt.Println("Enter the size of the grid (X Y):")
		str := term.ReadString()

		// Parse Grid Input
		pos, err := pkg.ParseGridSizeParams(str)
		if err != nil {
			terminal.Clear()
			term.RenderError(err.Error())
			continue
		}

		// Create Grid
		sessionGrid, err = grid.New(pos.X, pos.Y)
		if err != nil {
			terminal.Clear()
			term.RenderError(err.Error())
			continue
		}

		break
	}

	terminal.Clear()

	for {
		var d *pkg.AddRoverInput

		for {
			fmt.Println("Enter the rover's initial position and direction (X Y Direction):")
			str := term.ReadString()

			// Parse Grid Object Input
			details, err := pkg.ParseAddRoverCommands(str)
			if err != nil {
				terminal.Clear()
				term.RenderError(err.Error())
				continue
			}

			// Add Rover
			_, errs := sessionGrid.AddRover(details.X, details.Y, details.Dir)
			if len(errs) > 0 {
				terminal.Clear()
				term.RenderMultiError(errs)
				continue
			}
			d = details
			break
		}

		terminal.Clear()

		rover := sessionGrid.FindObjectByCoordinates(d.X, d.Y)

		fmt.Println("Enter the rover's movement commands:")
		str := term.ReadString()

		// Parse Rover Move Commands
		moveCommands, err := pkg.ParseRoverMoveCommands(str)
		if err != nil {
			terminal.Clear()
			term.RenderError(err.Error())
			continue
		}

		for _, v := range moveCommands {
			switch v {
			case "L":
				err = rover.RotateLeft()
				if err != nil {
					terminal.Clear()
					term.RenderError(err.Error())
					continue
				}
			case "R":
				err = rover.RotateRight()
				if err != nil {
					terminal.Clear()
					term.RenderError(err.Error())
					continue
				}
			case "M":
				switch *rover.Direction {
				case grid_object.North:
					errs := sessionGrid.IsValidPosition(rover.Position.X, rover.Position.Y+1)
					if len(errs) > 0 {
						terminal.Clear()
						term.RenderMultiError(errs)
						continue
					}
					err = rover.MoveNorth()
					if err != nil {
						terminal.Clear()
						term.RenderError(err.Error())
						continue
					}
				case grid_object.South:
					errs := sessionGrid.IsValidPosition(rover.Position.X, rover.Position.Y-1)
					if len(errs) > 0 {
						terminal.Clear()
						term.RenderMultiError(errs)
						continue
					}
					err = rover.MoveSouth()
					if err != nil {
						terminal.Clear()
						term.RenderError(err.Error())
						continue
					}
				case grid_object.East:
					errs := sessionGrid.IsValidPosition(rover.Position.X+1, rover.Position.Y)
					if len(errs) > 0 {
						terminal.Clear()
						term.RenderMultiError(errs)
						continue
					}
					err = rover.MoveEast()
					if err != nil {
						terminal.Clear()
						term.RenderError(err.Error())
						continue
					}
				case grid_object.West:
					errs := sessionGrid.IsValidPosition(rover.Position.X-1, rover.Position.Y)
					if len(errs) > 0 {
						terminal.Clear()
						term.RenderMultiError(errs)
						continue
					}
					err = rover.MoveWest()
					if err != nil {
						terminal.Clear()
						term.RenderError(err.Error())
						continue
					}
				}
			}
		}

		terminal.Clear()

		fmt.Println("Do you want to add another rover? (y/n)")
		str = term.ReadString()

		if str == "n" {
			break
		}
	}

	terminal.Clear()
	fmt.Println("Final positions of the rovers:")
	for i, rover := range sessionGrid.Objects {
		fmt.Printf("Rover %d: %d %d %s\n", i, rover.Position.X, rover.Position.Y, rover.Direction)
	}
}
