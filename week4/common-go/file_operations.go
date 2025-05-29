package main

import (
    "fmt"
    "os"
)

func main() {
    // Writing to a file
    data := []byte("Hello, Go File Operations!\nThis is a test file.")
    err := os.WriteFile("test.txt", data, 0644)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }

    // Reading from a file
    content, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    fmt.Println("File contents:", string(content))

    // Reading file info
    fileInfo, err := os.Stat("test.txt")
    if err != nil {
        fmt.Println("Error getting file info:", err)
        return
    }
    fmt.Printf("\nFile Info:\nName: %s\nSize: %d bytes\nModified: %v\n",
        fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime())
}
