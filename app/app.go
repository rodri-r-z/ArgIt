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
	error2 "github.com/rodri-r-z/argit/parser/error"
)

type App struct {
	Name             string
	Description      string
	Flags            map[string]*Flag
	Commands         map[string]*Command
	HelpMessage      string
	AutoGenerateHelp bool
}

func NewAp() *App {
	app := App{
		Flags:            make(map[string]*Flag),
		Commands:         make(map[string]*Command),
		AutoGenerateHelp: true,
	}

	// Add a help command
	app.AddFlag("help").
		AddAlias("h").
		SetDescription("shows this menu")

	return &app
}

func (app *App) SetName(val string) *App {
	app.Name = val
	return app
}

func (app *App) SetDescription(val string) *App {
	app.Description = val
	return app
}

func (app *App) SetAutoGenerateHelp(val bool) *App {
	app.AutoGenerateHelp = val
	return app
}

func (app *App) SetHelpMessage(val string) *App {
	app.HelpMessage = val
	return app
}

func (app *App) AddCommand(name string) *Command {
	// Make sure the map is initialized
	if app.Commands == nil {
		app.Commands = make(map[string]*Command)
	}

	// Create a new command
	cmd := Command{
		AutoGenerateUsage: true,
	}

	// Set the app
	cmd.app = app

	// Add a help flag
	cmd.AddFlag("help").
		AddAlias("h").
		SetDescription("shows this menu")

	// Make sure the command has a description
	if cmd.Description == "" {
		cmd.SetDescription("(description not set)")
	}

	// Save the command
	cmd.OriginalName = name
	app.Commands[name] = &cmd
	return &cmd
}

func (app *App) AddFlag(name string) *Flag {
	// Make sure the map is initialized
	if app.Flags == nil {
		app.Flags = make(map[string]*Flag)
	}

	// Create a new flag
	flag := NewFlag()

	// Save the flag
	flag.app = app
	flag.OriginalName = name
	app.Flags[name] = flag
	return flag
}

func (app *App) generateHelpImpl(err *error2.ArgvError) string {
	if !app.AutoGenerateHelp {
		return app.HelpMessage
	}

	// Check if the error comes from a command
	if err != nil {
		cmd, _ := err.SourceCommand.(*Command)
		if cmd != nil {
			return cmd.GenerateHelpWithError(err)
		}
	}

	return generateHelp(app, err, true, "", &app.Flags)
}

func (app *App) GenerateHelp() string {
	return app.generateHelpImpl(nil)
}

func (app *App) GenerateHelpWithError(err *error2.ArgvError) string {
	return app.generateHelpImpl(err)
}
