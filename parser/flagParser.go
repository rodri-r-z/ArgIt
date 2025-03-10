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
	"com.github.rodri-r-z/argit/app"
	"com.github.rodri-r-z/argit/parser/args"
	"com.github.rodri-r-z/argit/parser/command"
	error2 "com.github.rodri-r-z/argit/parser/error"
	"com.github.rodri-r-z/argit/parser/flag"
	"com.github.rodri-r-z/argit/types"
	"fmt"
)

func parseFlag(
	child *string,
	expectingCommandValue bool,
	cmd *app.Command,
	parsedCommand *command.ParsedCommand,
	source *app.App,
	result *args.Argv,
	lastFlag *flag.ParsedFlag,
	lastFlagName *string,
	expectingFlagValue *bool,
) *error2.ArgvError {
	childLen := len(*child)

	var traceCmd *app.Command
	if expectingCommandValue {
		traceCmd = cmd
	}

	// Ensure there are enough characters
	if childLen < 2 {
		return &error2.ArgvError{
			Code:          error2.FlagExpected,
			Message:       "Expected a flag name, none provided",
			SourceCommand: traceCmd,
		}
	}

	var flagName string
	// Check for long flags
	if (*child)[1] == '-' {
		if childLen < 3 {
			return &error2.ArgvError{
				Code:          error2.FlagExpected,
				Message:       "Expected a flag name, none provided",
				SourceCommand: traceCmd,
			}
		}

		flagName = (*child)[2:]
	} else {
		flagName = (*child)[1:]
	}

	var lookInto *map[string]*app.Flag
	var pushTo *map[string]*flag.ParsedFlag

	// Update the maps accordingly
	if expectingCommandValue {
		lookInto = &cmd.Flags
		pushTo = &parsedCommand.Flags
	} else {
		lookInto = &source.Flags
		pushTo = &result.Flags
	}

	// Retrieve the flag
	retrievedFlag := (*lookInto)[flagName]

	// Make sure the flag exists
	if retrievedFlag == nil {
		return &error2.ArgvError{
			Code:          error2.NoSuchFlag,
			Message:       fmt.Sprintf("No such flag %s", flagName),
			SourceCommand: traceCmd,
		}
	}

	newFlag := flag.ParsedFlag{
		Type: retrievedFlag.Type,
	}
	(*pushTo)[retrievedFlag.OriginalName] = &newFlag
	*lastFlag = newFlag
	*lastFlagName = retrievedFlag.OriginalName
	*expectingFlagValue = retrievedFlag.Type != types.Static

	return nil
}
