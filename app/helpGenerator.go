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
	error2 "com.github.rodri-r-z/argit/parser/error"
	"com.github.rodri-r-z/argit/types"
	"fmt"
	"strings"
)

func writeType(expected types.ArgType, builder *strings.Builder) {
	if expected != types.Static {
		builder.WriteString(" ")
		switch expected {
		case types.Bool:
			builder.WriteString("<true|false>")
		case types.String:
			builder.WriteString("<STRING>")
		case types.Float:
			builder.WriteString("<FLOAT>")
		case types.Int:
			builder.WriteString("<INTEGER>")
		default:
		}
	}
}

func writeCollection(
	name string,
	desc string,
	aliases []*string,
	writeNewLine bool,
	builder *strings.Builder,
	expected types.ArgType,
	autoGenerateUsage bool,
	usage string,
	required bool,
	isFlag bool,
) {
	if !autoGenerateUsage {
		builder.WriteString(usage)
	} else {
		// Build the string
		tempBuilder := strings.Builder{}

		if isFlag {
			tempBuilder.WriteString("--")
		}
		tempBuilder.WriteString(name)

		aliasesLen := len(aliases) - 1
		for i, el := range aliases {
			// Write an ", " where needed
			if i != aliasesLen || i == 0 {
				tempBuilder.WriteString(", ")
			}

			if isFlag {
				tempBuilder.WriteString("-")
			}
			tempBuilder.WriteString(*el)
		}

		// Write the expected type
		writeType(expected, &tempBuilder)

		builder.WriteString(
			fmt.Sprintf("%-35s %s", tempBuilder.String(), desc),
		)

		// Write if the flag is required
		if isFlag && required {
			builder.WriteString(" (REQUIRED)")
		}
	}

	// Write a newline if needed
	if writeNewLine {
		builder.WriteString("\n")
	}
}

func getMetadata(val any) (string, string, string, types.ArgType, []*string, bool, bool) {
	switch v := val.(type) {
	case *Command:
		return v.OriginalName, v.Description, v.Usage, v.Type, v.Aliases, true, v.AutoGenerateUsage
	case *Flag:
		return v.OriginalName, v.Description, v.Usage, v.Type, v.Aliases, v.Required, v.AutoGenerateUsage
	default:
		return "", "", "", types.Static, make([]*string, 0), false, false
	}
}

func collectAndWriteRawCollection[T *Command | *Flag](
	collection *map[string]T,
	builder *strings.Builder,
	isFlag bool,
) {
	// Save in a map the seen elements
	seen := make(map[T]bool)
	collectedLen := len(*collection) - 1
	i := 0
	for _, value := range *collection {
		if seen[value] {
			continue
		}

		seen[value] = true
		originalName, description, usage, argType, aliases, required, autoGenerateUsage := getMetadata(value)

		writeCollection(
			originalName,
			description,
			aliases,
			i != collectedLen,
			builder,
			argType,
			autoGenerateUsage,
			usage,
			required,
			isFlag,
		)
	}
}

func generateHelp(
	app *App,
	err *error2.ArgvError,
	printCommands bool,
	usageArgs string,
	flags *map[string]*Flag,
) string {
	// Use a strings.Builder
	builder := strings.Builder{}
	builder.WriteString(app.Name)
	builder.WriteString(" - ")
	builder.WriteString(app.Description)
	builder.WriteString("\n\n")

	// Print the error is needed
	if err != nil {
		builder.WriteString("Error: ")
		builder.WriteString(err.Message)
		builder.WriteString("\n\n")
	}

	// Print usage
	builder.WriteString("Usage: ")
	builder.WriteString(app.Name)
	builder.WriteString(" ")
	if usageArgs != "" {
		builder.WriteString(usageArgs)
	} else {
		builder.WriteString("[--global-flags...] <command> [--flags] <value> [--global-flags...]")
	}
	builder.WriteString("\nWhere:\n")
	builder.WriteString("- Values wrapped in <> are required\n")
	builder.WriteString("- Values wrapped in [] are not required")

	// Print the commands if needed
	if printCommands {
		builder.WriteString("\n\nAvailable commands:\n")
		// Write all commands
		collectAndWriteRawCollection(&app.Commands, &builder, false)
		builder.WriteString("\n")
	} else {
		builder.WriteString("\n\n")
	}

	builder.WriteString("Available flags:\n")
	// Write all flags
	collectAndWriteRawCollection(flags, &builder, true)

	return builder.String()
}
