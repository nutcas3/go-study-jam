package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run cmd/main.go")
		return
	}

	fmt.Println("Welcome to calculator\n\nSpecifying the operand to be used for operations. Choose between 'add', 'sub', 'mul' or 'div'.")

	var operation, input1, input2 string

	fmt.Scan(&operation)

	if strings.ToLower(operation) != "add" && strings.ToLower(operation) != "sub" && strings.ToLower(operation) != "mul" && strings.ToLower(operation) != "div" {
		fmt.Println("Invalid sign: Pick from 'add', 'sub', 'mul' or 'div'")
		return
	}

	fmt.Printf("You chose %q for the arithemtic sign. Input your first number\n", operation)
	fmt.Scan(&input1)

	fmt.Printf("You chose %q as your first number. Input your second number\n", input1)
	fmt.Scan(&input2)

	fmt.Printf("You chose %q as your second number.\n", input2)

	num1, err1 := strconv.ParseFloat(input1, 64)
	checkError(err1)

	num2, err2 := strconv.ParseFloat(input2, 64)
	checkError(err2)

	var result float64

	switch strings.ToLower(operation) {
	case "add":
		result = add(num1, num2)
	case "sub":
		result = subtract(num1, num2)
	case "mul":
		result = multiply(num1, num2)
	case "div":
		result = divide(num1, num2)
	default:
		fmt.Println("Unknown operation.")
		return
	}
	fmt.Printf("\nResult: %v\n", result)
}

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func divide(num1, num2 float64) float64 {
	if num2 == 0 {
		fmt.Println("Cannot divide by zero.")
		return 0.0
	}
	return num1 / num2
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Invalid numbers.")
		return
	}
}
