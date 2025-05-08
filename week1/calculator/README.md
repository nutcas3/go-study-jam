# Simple Calculator CLI

This project is part of Week 1 of the Go course. It is a simple command-line calculator application to practice basic Go syntax, variables, data types, operators, and control flow.

## How to run

Run the calculator from the `cmd` directory with the desired operation and numbers:

```bash
# Syntax:
go run cmd/main.go <add|sub|mul|div> <num1> <num2>

# Examples:
go run cmd/main.go add 2 3
go run cmd/main.go sub 10 4
go run cmd/main.go mul 5 6
go run cmd/main.go div 8 2
```

## Tasks
- Implement addition, subtraction, multiplication, and division.
- Use if/else and switch statements for operation selection.
- Handle invalid input gracefully.
