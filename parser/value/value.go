/*
   The ArgIt Project
   -----------------------------------------------------
   This code is released under the GNU GPL v3 license.
   For more information, please visit:
   https://www.gnu.org/licenses/gpl-3.0.html
   -----------------------------------------------------
   Copyright (c) 2025 Rodrigo R. & All Contributors
   This program comes with ABSOLUTELY NO WARRANTY.
   For details type `show w`. This is free software,
   and you are welcome to redistribute it under certain
   conditions; type `show c` for details.
*/

package value

import (
	"com.github.rodri-r-z/argit/app"
	error2 "com.github.rodri-r-z/argit/parser/error"
	"com.github.rodri-r-z/argit/types"
	"fmt"
	"strconv"
)

func ParseValue(value *string, expected types.ArgType, cmd *app.Command) (interface{}, *error2.ArgvError) {
	// Check for type mismatches
	switch expected {
	case types.Bool:
		if *value != "true" && *value != "false" {
			return nil, &error2.ArgvError{
				Code:          error2.TypeMismatch,
				Message:       fmt.Sprintf("I expected true|false, got %s", *value),
				SourceCommand: cmd,
			}
		}

		return *value == "true", nil
	case types.Int:
		// Attempt to parse the value
		parsedInt, err := strconv.ParseInt(*value, 10, 64)

		if err != nil {
			return nil, &error2.ArgvError{
				Code:          error2.TypeMismatch,
				Message:       fmt.Sprintf("I expected int, got %s", *value),
				SourceCommand: cmd,
			}
		}

		return parsedInt, nil
	case types.Float:
		// Attempt to parse the value
		parsedFloat, err := strconv.ParseFloat(*value, 64)

		if err != nil {
			return nil, &error2.ArgvError{
				Code:          error2.TypeMismatch,
				Message:       fmt.Sprintf("I expected float, got %s", *value),
				SourceCommand: cmd,
			}
		}

		return parsedFloat, nil
	case types.String:
		return *value, nil
	default:
	}

	return nil, nil
}
