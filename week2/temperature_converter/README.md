# Temperature Converter CLI

This project is part of Week 2 of the Go course. It is a command-line application for converting temperatures between Celsius and Fahrenheit.

## How to run

Run the temperature converter from the `cmd` directory with the desired mode and temperature:

```bash
# Syntax:
go run cmd/main.go <CtoF|FtoC> <temperature>

# Examples:
go run cmd/main.go CtoF 100
# Converts 100°C to Fahrenheit

go run cmd/main.go FtoC 212
# Converts 212°F to Celsius
```

## Tasks
- Accept user input for temperature and unit.
- Convert between Celsius and Fahrenheit.
- Practice functions, error handling, and input validation.
