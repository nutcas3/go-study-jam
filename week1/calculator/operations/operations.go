package operations

import "fmt"

func Add(num1, num2 float64) (float64, error) {
	return num1 + num2, nil
}

func Subtract(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}

func Multiply(num1, num2 float64) (float64, error) {
	return num1 * num2, nil
}

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}

func AddIterable(nums []float64) (float64, error) {
	var sum float64

	for _, num := range nums {
		sum += num
	}
	return sum, nil
}
