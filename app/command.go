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

package app

import (
	"fmt"
	error2 "github.com/rodri-r-z/argit/parser/error"
	"github.com/rodri-r-z/argit/types"
	"strings"
)

type Command struct {
	app               *App
	Description       string
	Usage             string
	AutoGenerateUsage bool
	Flags             map[string]*Flag
	Aliases           []*string
	Type              types.ArgType
	OriginalName      string
}

func (cmd *Command) SetDescription(val string) *Command {
	cmd.Description = val
	return cmd
}

func (cmd *Command) SetType(val types.ArgType) *Command {
	cmd.Type = val
	return cmd
}

func (cmd *Command) SetAutoGenerateUsage(val bool) *Command {
	cmd.AutoGenerateUsage = val
	return cmd
}

func (cmd *Command) SetUsage(val string) *Command {
	cmd.Usage = val
	return cmd
}

func (cmd *Command) AddFlag(name string) *Flag {
	// Make sure the map is initialized
	if cmd.Flags == nil {
		cmd.Flags = make(map[string]*Flag)
	}

	// Create a new flag
	flag := NewFlag()

	// Check if we already have a flag that has this name
	if cmd.Flags[name] != nil {
		panic(fmt.Sprintf("a flag named '%s' is already registered", name))
	}

	// Save the flag
	flag.command = cmd
	flag.OriginalName = name
	cmd.Flags[name] = flag
	return flag
}

func (cmd *Command) AddAlias(name string) *Command {
	// Check if app is nil
	if cmd.app == nil {
		panic("app is nil (did you call app.AddCommand first?)")
	}

	// Check if we already have a command that has this name
	if cmd.app.Commands[name] != nil {
		panic(fmt.Sprintf("a command named '%s' is already registered", name))
	}

	// Directly add the same command to the app for better lookup
	cmd.Aliases = append(cmd.Aliases, &name)
	cmd.app.Commands[name] = cmd
	return cmd
}

func (cmd *Command) generateUsageArgs() string {
	// Use a strings.Builder for efficiency
	builder := strings.Builder{}

	builder.WriteString("[--global-flags...] <")
	builder.WriteString(cmd.OriginalName)

	// Iterate the app's commands to find out aliases
	for _, name := range cmd.Aliases {
		// Write the alias
		builder.WriteString("|")
		builder.WriteString(*name)
	}

	builder.WriteString(">")

	// Write the command's type
	writeType(cmd.Type, &builder)
	builder.WriteString(" [--flags] [--global-flags...]")

	return builder.String()
}

func (cmd *Command) GenerateHelpWithError(err *error2.ArgvError) string {
	if !cmd.AutoGenerateUsage {
		return cmd.Usage
	}

	return generateHelp(cmd.app, err, false, cmd.generateUsageArgs(), &cmd.Flags)
}
