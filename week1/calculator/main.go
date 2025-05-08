package main

import "fmt"

func main() {
    // Basic types
    var integer int = 10
    var floatingPoint float64 = 3.14
    var complexNumber complex128 = complex(1, 2)
    var text string = "Hello, Go!"
    var boolean bool = true

    fmt.Println("Integer:", integer)
    fmt.Println("Floating-point:", floatingPoint)
    fmt.Println("Complex number:", complexNumber)
    fmt.Println("String:", text)
    fmt.Println("Boolean:", boolean)

    // Aggregate types
    var array [3]int = [3]int{1, 2, 3}
    type Person struct {
        Name string
        Age  int
    }
    var person Person = Person{"Alice", 30}

    fmt.Println("Array:", array)
    fmt.Println("Struct:", person)

    // Reference types
    var pointer *int = &integer
    var slice []int = []int{4, 5, 6}
    var mapData map[string]int = map[string]int{"one": 1, "two": 2}

    fmt.Println("Pointer:", pointer)
    fmt.Println("Slice:", slice)
    fmt.Println("Map:", mapData)
}