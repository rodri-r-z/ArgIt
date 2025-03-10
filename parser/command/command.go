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

package command

import (
	flag2 "com.github.rodri-r-z/argit/parser/converter/flag"
	"com.github.rodri-r-z/argit/parser/converter/value"
	"com.github.rodri-r-z/argit/parser/flag"
)

type ParsedCommand struct {
	Value  *interface{}
	Flags  map[string]*flag.ParsedFlag
	Source interface{}
}

func (c *ParsedCommand) Int() int {
	return value.ConvertInt(c.Value)
}

func (c *ParsedCommand) Float() float64 {
	return value.ConvertFloat(c.Value)
}

func (c *ParsedCommand) Bool() bool {
	return value.ConvertBool(c.Value)
}

func (c *ParsedCommand) String() string {
	return value.ConvertString(c.Value)
}

func (c *ParsedCommand) GetString(name string) string {
	return flag2.GetStringFlag(&c.Flags, name)
}

func (c *ParsedCommand) GetInt(name string) int {
	return flag2.GetIntFlag(&c.Flags, name)
}

func (c *ParsedCommand) GetFloat(name string) float64 {
	return flag2.GetFloatFlag(&c.Flags, name)
}

func (c *ParsedCommand) GetBool(name string) bool {
	return flag2.GetBoolFlag(&c.Flags, name)
}

func (c *ParsedCommand) HasFlag(name string) bool {
	// Get the flag
	f := c.Flags[name]
	return f != nil
}
