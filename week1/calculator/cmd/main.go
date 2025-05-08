package main

import (
	"fmt"
	"os"
	"strconv"
	"calculator/macus"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <add|sub|mul|div> <num1> <num2>")
		return
	}
	operation := os.Args[1]
	num1, err1 := strconv.ParseFloat(os.Args[2], 64)
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers.")
		return
	}
	var result float64
	// rbac
	switch operation {
	case "add":
		result = num1 + num2
	case "sub":
		result = num1 - num2
	case "mul":
		result = num1 * num2
	case "div":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Unknown operation.")
		return
	}
	fmt.Printf("Result: %v\n", result)

	// Show Go language examples
	kendi.ShowAllExamples()
}
