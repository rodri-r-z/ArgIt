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

package flag

import (
	"com.github.rodri-r-z/argit/parser/converter/value"
	"com.github.rodri-r-z/argit/types"
)

type ParsedFlag struct {
	Type  types.ArgType
	Value *interface{}
}

func (f *ParsedFlag) Int() int {
	return value.ConvertInt(f.Value)
}

func (f *ParsedFlag) Float() float64 {
	return value.ConvertFloat(f.Value)
}

func (f *ParsedFlag) Bool() bool {
	return value.ConvertBool(f.Value)
}

func (f *ParsedFlag) String() string {
	return value.ConvertString(f.Value)
}
