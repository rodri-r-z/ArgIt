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

package error

type ArgvErrorCode int

const (
	NoParams ArgvErrorCode = iota
	NoCommands
	NoSuchCommand
	NoSuchFlag
	CommandIsStatic
	CommandMustHaveValue
	FlagMissing
	FlagMustHaveValue
	ValueNotExpected
	FlagExpected
	TypeMismatch
)

type ArgvError struct {
	Code          ArgvErrorCode
	Message       string
	SourceCommand interface{}
}

func (e *ArgvError) Panic() {
	panic(e.Message)
}
