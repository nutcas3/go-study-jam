package kendi

import (
	"fmt"
	"time"
)

// ShowControlFlow demonstrates Go's control flow structures
func ShowControlFlow() {
	fmt.Println("\n--- Go Control Flow Examples ---")

	// 1. If-Else statements
	fmt.Println("\n1. If-Else Statements:")
	
	x := 10
	if x > 5 {
		fmt.Println("  x is greater than 5")
	} else if x < 5 {
		fmt.Println("  x is less than 5")
	} else {
		fmt.Println("  x is equal to 5")
	}

	// If with a short statement
	if y := 15; y > x {
		fmt.Println("  y is greater than x")
	}

	// 2. For loops
	fmt.Println("\n2. For Loops:")
	
	// Standard C-like for loop
	fmt.Println("  Standard for loop:")
	for i := 0; i < 3; i++ {
		fmt.Printf("    i = %d\n", i)
	}

	// For as a while loop
	fmt.Println("  For as while loop:")
	j := 0
	for j < 3 {
		fmt.Printf("    j = %d\n", j)
		j++
	}

	// Infinite loop with break
	fmt.Println("  For with break:")
	k := 0
	for {
		fmt.Printf("    k = %d\n", k)
		k++
		if k >= 3 {
			break
		}
	}

	// For-range loop (iterating over collections)
	fmt.Println("  For-range loop:")
	colors := []string{"red", "green", "blue"}
	for index, color := range colors {
		fmt.Printf("    colors[%d] = %s\n", index, color)
	}

	// 3. Switch statements
	fmt.Println("\n3. Switch Statements:")
	
	day := time.Now().Weekday()
	fmt.Printf("  Today is %s\n", day)
	
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("  It's the weekend!")
	default:
		fmt.Println("  It's a weekday.")
	}

	// Switch without an expression (alternative to if-else chains)
	fmt.Println("  Switch without expression:")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("    Good morning!")
	case t.Hour() < 17:
		fmt.Println("    Good afternoon!")
	default:
		fmt.Println("    Good evening!")
	}

	// 4. Defer
	fmt.Println("\n4. Defer Statement:")
	fmt.Println("  Defer executes functions in LIFO order after surrounding function returns")
	
	defer fmt.Println("    This prints last (3)")
	defer fmt.Println("    This prints second (2)")
	fmt.Println("    This prints first (1)")

	// 5. Panic and Recover (commented out to avoid actual panic)
	fmt.Println("\n5. Panic and Recover:")
	fmt.Println("  Panic is for exceptional errors, recover can catch panics")
	fmt.Println("  Example: func() { defer recoverFunc(); panic(\"something bad happened\") }")

	// 6. Goto (rarely used)
	fmt.Println("\n6. Goto Statement:")
	fmt.Println("  Goto is available but rarely used in Go")
	
	count := 0
loop:
	fmt.Printf("  count = %d\n", count)
	count++
	if count < 3 {
		goto loop
	}

	fmt.Println("\n--- End of Control Flow Examples ---")
}
