<div align="center">
    <h1>The ArgIt Project</h1>
    ArgIt is a blazingly fast, memory-efficient
    and developer-friendly CLI arguments parser
</div>

---

## 👋 Welcome

Welcome to the official ArgIt repository! In this file, you'll find a guide on how to style your code
to contribute to the ArgIt Project.

---

## 📝 The basics

Please note that any code that does not follow the guidelines in this document will not be accepted.

**Guidelines:**

- Tab for indentation
- No trailing whitespace or empty lines at the end of files
- No outrageously long lines of code, try to keep logic as concise as possible
  - For example:
    - ```go
      // Bad
      if err != nil { 
        return err
      }
        ```
      
    - ```go
      // Bad
	  if err != nil || my_file_is_okay > 0 && ((myOtherFile == "okay" || myOtherFile == "not okay") || fetchSomeResource() == "13.5") {
        return err
	  }
        ```
      
- No commented-out code, if you remove code, you must provide a comment explaining why it is no longer needed
- No unused variables or imports
- No unnecessary type casting
- No syntax errors

**1. Function invocations**

Unless the parameters are too long, function invocations should be on the same line.

Example:

```go
// Good
myFunction(param1, param2, param3)

// Also Good
myFunction(
    myOtherFunction(param1).accessSomething(),
	myOtherFunction(param2).accessSomething(),
    myOtherFunction(param3).accessSomething(),
)

// Bad
myFunction(
    param1,
    param2,
    param3,
)
```

**3. Comments**

Avoid block comments, use single-line comments instead.

Example:

```go
// Good
// This is a comment

// Bad
/*
This is a comment
*/
```

**4. Naming conventions**

- Use `camelCase` for variables and functions
- Use `PascalCase` for types and interfaces
- Use `UPPER_SNAKE_CASE` for constants
- Use `camelCase` for file and directory names
- Use lowercase for package names

**5. Error handling**

Always handle errors, do not ignore them.

Example:

```go
// Good
if err != nil {
    return err
}

// Bad
if err != nil {
    // Do nothing
}
```

Unless there is underlying logic that requires it, do not use `panic()`.
Instead use the `logger` package to log information, warns and errors.

**6. Copyright**

All files must contain the following header at the top of the file:

```go
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
```