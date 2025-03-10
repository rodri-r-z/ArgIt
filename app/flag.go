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
	"com.github.rodri-r-z/argit/types"
	"fmt"
)

type Flag struct {
	app               *App
	command           *Command
	Usage             string
	Description       string
	AutoGenerateUsage bool
	Type              types.ArgType
	Aliases           []*string
	Required          bool
	OriginalName      string
}

func NewFlag() *Flag {
	return &Flag{
		AutoGenerateUsage: true,
		Description:       "(description not set)",
	}
}

func (flag *Flag) SetAutoGenerateUsage(val bool) *Flag {
	flag.AutoGenerateUsage = val
	return flag
}

func (flag *Flag) SetUsage(val string) *Flag {
	flag.Usage = val
	return flag
}

func (flag *Flag) SetDescription(val string) *Flag {
	flag.Description = val
	return flag
}

func (flag *Flag) SetType(val types.ArgType) *Flag {
	flag.Type = val
	return flag
}

func (flag *Flag) SetRequired(val bool) *Flag {
	flag.Required = val
	return flag
}

func (flag *Flag) AddAlias(name string) *Flag {
	// Check if app and command are nil
	if flag.command == nil && flag.app == nil {
		panic("command is nil (did you call app.AddFlag/cmd.AddFlag first?)")
	}

	var flags *map[string]*Flag

	// Update the flags variable accordingly
	if flag.command != nil {
		flags = &flag.command.Flags
	} else {
		flags = &flag.app.Flags
	}

	// Check if we already have a command that has this name
	if (*flags)[name] != nil {
		panic(fmt.Sprintf("a flag named '%s' is already registered", name))
	}

	// Directly add the same command to the app for better lookup
	flag.Aliases = append(flag.Aliases, &name)
	(*flags)[name] = flag

	return flag
}
