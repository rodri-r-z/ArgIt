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

package parser

import (
	"fmt"
	"github.com/rodri-r-z/argit/app"
	error2 "github.com/rodri-r-z/argit/parser/error"
	"github.com/rodri-r-z/argit/parser/flag"
	value2 "github.com/rodri-r-z/argit/parser/value"
)

func parseFlagValue(
	isFlag bool,
	child *string,
	lastFlag *flag.ParsedFlag,
	lastFlagName string,
	cmd *app.Command,
) *error2.ArgvError {
	if isFlag {
		return &error2.ArgvError{
			Code:          error2.FlagMustHaveValue,
			Message:       fmt.Sprintf("Expected a value for flag %s, flag provided", lastFlagName),
			SourceCommand: cmd,
		}
	}

	if lastFlag == nil {
		panic("Failed to parse arguments - Last parsed flag is nil. Please report this incident.")
	}

	// Parse the value
	value, err := value2.ParseValue(child, lastFlag.Type, cmd)
	if err != nil {
		return err
	}

	lastFlag.Value = &value

	return nil
}
