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

package args

import (
	"github.com/rodri-r-z/argit/parser/command"
	flag2 "github.com/rodri-r-z/argit/parser/converter/flag"
	"github.com/rodri-r-z/argit/parser/flag"
)

type Argv struct {
	ExecPath string
	Argc     int
	Flags    map[string]*flag.ParsedFlag
	Command  *command.ParsedCommand
}

func (a *Argv) GetBool(name string) bool {
	return flag2.GetBoolFlag(&a.Flags, name)
}

func (a *Argv) GetString(name string) string {
	return flag2.GetStringFlag(&a.Flags, name)
}

func (a *Argv) GetInt(name string) int {
	return flag2.GetIntFlag(&a.Flags, name)
}

func (a *Argv) GetFloat(name string) float64 {
	return flag2.GetFloatFlag(&a.Flags, name)
}

func (a *Argv) HasFlag(name string) bool {
	// Get the flag
	f := a.Flags[name]
	return f != nil
}
