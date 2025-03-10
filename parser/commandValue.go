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
	"github.com/rodri-r-z/argit/app"
	"github.com/rodri-r-z/argit/parser/command"
	error2 "github.com/rodri-r-z/argit/parser/error"
	value2 "github.com/rodri-r-z/argit/parser/value"
)

func parseCommandValue(
	child *string,
	cmd *app.Command,
	parsedCommand *command.ParsedCommand,
) *error2.ArgvError {
	// Parse the value
	value, err := value2.ParseValue(child, cmd.Type, cmd)
	if err != nil {
		return err
	}

	parsedCommand.Value = &value
	return nil
}
