package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: tempconv <CtoF|FtoC> <temperature>")
		return
	}
	mode := os.Args[1]
	temp, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid temperature.")
		return
	}
	switch mode {
	case "CtoF":
		fmt.Printf("%.2f°C = %.2f°F\n", temp, temp*9/5+32)
	case "FtoC":
		fmt.Printf("%.2f°F = %.2f°C\n", temp, (temp-32)*5/9)
	default:
		fmt.Println("Unknown mode. Use CtoF or FtoC.")
	}
}
