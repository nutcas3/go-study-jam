package main

import (
	"calculator/operations"
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

	fmt.Println("Welcome to calculator.\nChoose an operation: 'add', 'sub', 'mul', or 'div'.")

	// The significance of the infinite loop at this point is to ensure the calculator runs even after a successful operation. Closes when the user exits.
	for {
		var operation string

		// The infinite loop at this point ensures the operation persists until a valid operation is provided.
		for {
			fmt.Print("\nEnter operation: ")
			fmt.Scan(&operation)
			operation = strings.ToLower(strings.TrimSpace(operation))

			if operation == "add" || operation == "sub" || operation == "mul" || operation == "div" {
				break
			}
			fmt.Println("Invalid operation. Choose from 'add', 'sub', 'mul', or 'div'.")
		}

		// Only proceeds to the next step if a valid number is provided by the user.
		var input1 string
		var num1 float64
		for {
			fmt.Print("Enter input 1: ")
			fmt.Scan(&input1)
			n1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
			if err == nil {
				num1 = n1
				break
			}
			fmt.Println("Invalid number. Try again.")
		}

		// Only proceeds to the next step if a valid number is provided by the user.
		var input2 string
		var num2 float64
		for {
			fmt.Print("Enter input 2: ")
			fmt.Scan(&input2)
			n2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
			if err == nil {
				num2 = n2
				break
			}
			fmt.Println("Invalid number. Try again.")
		}

		var result float64

		switch operation {
		case "add":
			result = operations.Add(num1, num2)
		case "sub":
			result = operations.Subtract(num1, num2)
		case "mul":
			result = operations.Multiply(num1, num2)
		case "div":
			result = operations.Divide(num1, num2)
		}
		fmt.Printf("Result: %v\n", result)
	}
}
