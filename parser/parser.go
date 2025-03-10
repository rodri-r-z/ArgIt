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
	"github.com/rodri-r-z/argit/parser/args"
	"github.com/rodri-r-z/argit/parser/command"
	error2 "github.com/rodri-r-z/argit/parser/error"
	"github.com/rodri-r-z/argit/parser/flag"
	"github.com/rodri-r-z/argit/types"
	"os"
)

func checkMissingFlags(
	flags *map[string]*flag.ParsedFlag,
	source *map[string]*app.Flag,
	cmd *app.Command,
) *error2.ArgvError {
	// Iterate over the flags
	for name, f := range *source {
		if !f.Required {
			continue
		}

		if (*flags)[name] == nil {
			return &error2.ArgvError{
				Code:          error2.FlagMissing,
				Message:       fmt.Sprintf("The flag '%s' is required, but it is missing", name),
				SourceCommand: cmd,
			}
		}
	}

	return nil
}

func ParseArgv(argv []string, source *app.App) (*args.Argv, *error2.ArgvError) {
	result := args.Argv{
		// Set the executable's path
		ExecPath: argv[0],
		Argc:     len(argv),
		Flags:    make(map[string]*flag.ParsedFlag),
	}

	// Check for empty arguments
	if result.Argc == 1 {
		if len(source.Commands) == 0 {
			return nil, &error2.ArgvError{
				Code:    error2.NoCommands,
				Message: "I expected at least one command in your app, found zero",
			}
		} else {
			return nil, &error2.ArgvError{
				Code:    error2.NoParams,
				Message: "I expected at least one command, found zero",
			}
		}
	}

	// Make sure we have a value for this command
	if result.Argc < 2 {
		return nil, &error2.ArgvError{
			Code:    error2.CommandMustHaveValue,
			Message: "Not enough arguments provided",
		}
	}

	// Create a ParsedCommand
	parsedCommand := command.ParsedCommand{
		Flags: make(map[string]*flag.ParsedFlag),
	}

	// Set the parsed command to the result
	result.Command = &parsedCommand

	// Determine if we are expecting a value
	expectingCommand := true
	expectingCommandValue := false
	expectingFlagValue := false
	var lastFlag *flag.ParsedFlag
	var lastFlagName string
	var cmdName string
	var cmd *app.Command

	for i := 1; i < result.Argc; i++ {
		child := argv[i]
		isFlag := child[0] == '-'

		// Parse flag values first
		if expectingFlagValue {
			var traceCmd *app.Command

			if expectingCommandValue {
				traceCmd = cmd
			}

			err := parseFlagValue(
				isFlag,
				&child,
				lastFlag,
				&expectingFlagValue,
				lastFlagName,
				traceCmd,
			)

			if err != nil {
				return nil, err
			}

			continue
		}

		// Parse commands
		if !isFlag && expectingCommand {
			// Retrieve the command
			gotCommand := source.Commands[child]

			// Make sure the command exists
			if gotCommand == nil {
				return nil, &error2.ArgvError{
					Code:          error2.NoSuchCommand,
					Message:       fmt.Sprintf("No such command %s", child),
					SourceCommand: cmd,
				}
			}

			// Check if the command is static
			if gotCommand.Type == types.Static {
				return nil, &error2.ArgvError{
					Code:          error2.CommandIsStatic,
					Message:       fmt.Sprintf("The command %s is static", child),
					SourceCommand: cmd,
				}
			}

			cmdName = child
			cmd = gotCommand
			parsedCommand.Source = cmd
			expectingCommand = false
			expectingCommandValue = true
			continue
		}

		// Parse flags
		if isFlag {
			err := parseFlag(
				&child,
				expectingCommandValue,
				cmd,
				&parsedCommand,
				source,
				&result,
				lastFlag,
				&lastFlagName,
				&expectingFlagValue,
			)

			if err != nil {
				return nil, err
			}

			continue
		}

		if expectingCommandValue {
			err := parseCommandValue(
				&child,
				cmd,
				&parsedCommand,
				&expectingFlagValue,
			)

			if err != nil {
				return nil, err
			}
			continue
		}

		return nil, &error2.ArgvError{
			Code:          error2.ValueNotExpected,
			Message:       fmt.Sprintf("Expected flag/command, value '%s' provided", child),
			SourceCommand: cmd,
		}
	}

	var traceCmd *app.Command

	if expectingCommandValue {
		traceCmd = cmd
	}

	if expectingFlagValue || expectingCommandValue {
		if lastFlag != nil && lastFlag.Type != types.Static {
			return nil, &error2.ArgvError{
				Code:          error2.FlagMustHaveValue,
				Message:       fmt.Sprintf("Flag %s must have a value, none provided", lastFlagName),
				SourceCommand: traceCmd,
			}
		} else if cmd != nil && lastFlagName != "help" {
			return nil, &error2.ArgvError{
				Code:          error2.CommandMustHaveValue,
				Message:       fmt.Sprintf("Command %s must have a value, none provided", cmdName),
				SourceCommand: traceCmd,
			}
		}
	}

	// Check for help flags in the source
	if result.HasFlag("help") {
		fmt.Print(source.GenerateHelp())
		os.Exit(0)
	}

	// Sanitize the command
	if result.Command != nil {
		// Assert the command
		cmd, _ := result.Command.Source.(*app.Command)

		// Check for help flags in the commands
		if result.Command.HasFlag("help") {
			fmt.Print(cmd.GenerateHelpWithError(nil))
			os.Exit(0)
		}

		// Check for missing flags
		err := checkMissingFlags(&result.Command.Flags, &cmd.Flags, traceCmd)
		if err != nil {
			return nil, err
		}
	}

	// Check for missing flags
	err := checkMissingFlags(&result.Flags, &source.Flags, nil)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ParseArgvFromOs(app *app.App) (*args.Argv, *error2.ArgvError) {
	return ParseArgv(os.Args, app)
}
