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
	"com.github.rodri-r-z/argit/parser/flag"
	"com.github.rodri-r-z/argit/types"
)

func processFlag[T comparable](f *flag.ParsedFlag, expected types.ArgType, fallback T) T {
	// Ensure the flag is not nil
	if f == nil {
		return fallback
	}

	// Match the types
	if f.Type != expected {
		return fallback
	}

	// Get the value
	return value.GetOrFallback(f.Value, fallback)
}

func GetStringFlag(flagMap *map[string]*flag.ParsedFlag, name string) string {
	return processFlag((*flagMap)[name], types.String, "")
}

func GetBoolFlag(flagMap *map[string]*flag.ParsedFlag, name string) bool {
	return processFlag((*flagMap)[name], types.Bool, false)
}

func GetIntFlag(flagMap *map[string]*flag.ParsedFlag, name string) int {
	return processFlag((*flagMap)[name], types.Int, 0)
}

func GetFloatFlag(flagMap *map[string]*flag.ParsedFlag, name string) float64 {
	return processFlag((*flagMap)[name], types.Float, 0.0)
}
