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

	fmt.Println("Welcome to calculator.\nChoose an operation: 'add', 'addIterable', 'sub', 'mul', or 'div'.\nNote: 'addIterable' allows you to add multiple numbers at once.")

	// The significance of the infinite loop at this point is to ensure the calculator runs even after a successful operation. Closes when the user exits.
	for {
		var operation, digitCount string

		// The infinite loop at this point ensures the operation persists until a valid operation is provided.
		for {
			fmt.Print("\nEnter operation: ")
			fmt.Scan(&operation)
			operation = strings.ToLower(strings.TrimSpace(operation))

			if operation == "add" || operation == "additerable" || operation == "sub" || operation == "mul" || operation == "div" {
				if operation == "additerable" {
					fmt.Print("How many numbers do you wish to input? ")
					fmt.Scan(&digitCount)

					count, err := strconv.Atoi(digitCount)
					checkError(err)

					num := make([]float64, 0, count)

					for i := 0; i < int(count); i++ {
						var input string

						fmt.Printf("Enter input %d: ", i+1)
						fmt.Scan(&input)

						input_, loopErr := strconv.ParseFloat(input, 64)
						checkError(loopErr)

						num = append(num, input_)
					}

					// Used continue intead of return for this block to let the calculator skip the rest of the opration.
					if len(num) == 0 {
						fmt.Println("No numbers were provided for addition.")
						continue
					}
					result, err := operations.AddIterable(num)
					checkError(err)
					fmt.Printf("Result: %v\n", result)
					continue
				}
				break
			}
			fmt.Println("Invalid operation. Choose from 'add', 'addIterable', 'sub', 'mul', or 'div'.")
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

		var (
			result float64
			err    error
		)

		switch operation {
		case "add":
			result, err = operations.Add(num1, num2)
			checkError(err)
		case "sub":
			result, err = operations.Subtract(num1, num2)
			checkError(err)
		case "mul":
			result, err = operations.Multiply(num1, num2)
			checkError(err)
		case "div":
			result, err = operations.Divide(num1, num2)
			checkError(err)
		}
		fmt.Printf("Result: %v\n", result)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Invalid number.")
		return
	}
}
