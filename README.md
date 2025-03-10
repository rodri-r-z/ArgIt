<div align="center">
    <h1>The ArgIt Project</h1>
    ArgIt is a blazingly fast, memory-efficient
    and developer-friendly CLI arguments parser
</div>

---

## 👋 Welcome

Welcome to the official ArgIt repository!
In this file, you'll find an introduction to
ArgIt.

## 📦 Features

- Blazing-fast parsing times
- Simple API
- **[Free - as in Freedom.](LICENSE)**
- And more!

## ❌ What we don't have (yet)

- Subcommands
- Colored terminal outputs

---

## 🚀 Getting Started

Getting started with ArgIt is super easy,
just make sure you have the library installed
by adding it in your `go.mod`

Example program:

```go
package main

import (
	app2 "com.github.rodri-r-z/argit/app"
	"com.github.rodri-r-z/argit/parser"
	"com.github.rodri-r-z/argit/types"
	"fmt"
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
```

---

## 📦 Contributions

Contributions are welcome! If you'd like to contribute to the Fluent language, please read the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information.

---

## 🔒 Security

Please refer to [SECURITY.md](SECURITY.md) for more information on how to report security vulnerabilities.

---

## 📝 License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for more information.

```
Copyright (C) 2024 Rodrigo R. & All Contributors
This program comes with ABSOLUTELY NO WARRANTY; for details type `show w`.
This is free software, and you are welcome to redistribute it under certain conditions;
type `show c` for details.
```