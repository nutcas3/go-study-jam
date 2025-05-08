package kendi

import "fmt"

// Person is a simple struct type for demonstration
type Person struct {
	Name string
	Age  int
}

// Greeter is an interface that defines a Greet method
type Greeter interface {
	Greet() string
}

// Robot implements the Greeter interface
type Robot struct{}

// Greet returns a robot greeting
func (r Robot) Greet() string { 
	return "beep boop" 
}

// ShowGoTypesExample demonstrates Go's main data types in action.
func ShowGoTypesExample() {
	fmt.Println("\n--- Go Data Types Example ---")

	// Basic types
	var i int = 42
	var f float64 = 3.14
	var s string = "hello"
	var b bool = true
	fmt.Println("int:", i, "float64:", f, "string:", s, "bool:", b)

	// Array (fixed size)
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("array:", arr)

	// Slice (dynamic size)
	slc := []string{"go", "is", "fun"}
	fmt.Println("slice:", slc)

	// Map (key-value)
	m := map[string]int{"alice": 25, "bob": 30}
	fmt.Println("map:", m)

	// Struct (custom type)
	p := Person{Name: "Charlie", Age: 28}
	fmt.Println("struct:", p)

	// Pointer
	var x int = 100
	var px *int = &x
	fmt.Println("pointer to x:", px, "value at pointer:", *px)

	// Function as variable
	add := func(a, b int) int { return a + b }
	fmt.Println("function as variable (add):", add(2, 3))

	// Interface
	var g Greeter = Robot{}
	fmt.Println("interface (Greeter):", g.Greet())

	// rune and byte
	var ch rune = 'G' // Unicode code point
	var by byte = 'A' // Raw byte
	fmt.Println("rune:", ch, "byte:", by)

	// error
	var myErr error = fmt.Errorf("something went wrong")
	fmt.Println("error:", myErr)

	// fmt.Println("--- End of Example ---\n")
}
