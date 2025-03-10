package main

import (
	"fmt"
	app2 "github.com/rodri-r-z/argit/app"
	"github.com/rodri-r-z/argit/parser"
	"github.com/rodri-r-z/argit/types"
)

func main() {
	app := app2.NewAp().
		SetName("my_cli").
		SetDescription("An example CLI")

	app.AddFlag("test").
		SetType(types.Bool)

	app.AddCommand("greet").
		SetDescription("greets  you").
		SetType(types.String).
		AddAlias("g").
		AddFlag("excited").
		SetType(types.Bool).
		AddAlias("e")

	argv, err := parser.ParseArgvFromOs(app)
	if err != nil {
		fmt.Print(app.GenerateHelpWithError(err))
		return
	}

	fmt.Print("Hello, ", argv.Command.String())
	if argv.Command.GetBool("excited") {
		fmt.Print("!!!!")
	}
	fmt.Println()

	for name, el := range argv.Flags {
		fmt.Printf("Flag %s, value: %v\n", name, el.Bool())
	}
}
