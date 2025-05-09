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

	fmt.Println("Welcome to calculator\n\nSpecifying the operand to be used for operations. Choose between 'add', 'addIterable','sub', 'mul' or 'div'.")
	fmt.Println("Note: 'addIterable' allows you to add multiple numbers at once.")
	var operation, input1, input2, digitCount string

	fmt.Scan(&operation)

	if strings.ToLower(operation) != "add" && strings.ToLower(operation) != "sub" && strings.ToLower(operation) != "mul" && strings.ToLower(operation) != "div" && strings.ToLower(operation) != "additerable" {
		fmt.Println("Invalid sign: Pick from 'add', 'addIterable','sub', 'mul' or 'div'")
		return
	}
    if strings.ToLower(operation) == "additerable" {
		fmt.Printf("How many numbers do you wish to input? ")
		fmt.Scan(&digitCount)
		count, err1 := strconv.Atoi(digitCount)
		checkError(err1)
		num := make([]float64, 0, count)
		for i := 0; i < int(count); i++ {
			var input string
			fmt.Printf("Enter input %d: ", i+1)
			fmt.Scan(&input)
			input_, loopErr := strconv.ParseFloat(input, 64)
			checkError(loopErr)
			num = append(num, input_)
		}
		if len(num) == 0 {
			fmt.Println("No numbers were provided for addition.")
			return	
		}
		result, err := addIterable(num)
		checkError(err)
		fmt.Printf("\nResult: %v\n", result)
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

func addIterable(nums []float64 ) (float64, error) {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Invalid numbers.")
		return
	}
}
