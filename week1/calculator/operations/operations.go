package operations

import "fmt"

func Add(num1, num2 float64) float64 {
	return num1 + num2
}

func Subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1, num2 float64) float64 {
	if num2 == 0 {
		fmt.Println("Cannot divide by zero.")
		return 0.0
	}
	return num1 / num2
}
