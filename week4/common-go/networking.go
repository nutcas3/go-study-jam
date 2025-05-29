package main

import (
    "fmt"
    "io"
    "net"
)

func main() {
    // Simple TCP server
    go tcpServer()

    // TCP client
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    // Send data
    fmt.Fprintf(conn, "Hello from client!")

    // Read response
    response, err := io.ReadAll(conn)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }
    fmt.Printf("Server response: %s\n", response)
}

func tcpServer() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    // Read incoming data
    data, err := io.ReadAll(conn)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }

    // Echo back
    fmt.Printf("Received: %s\n", data)
    conn.Write([]byte("Server received your message!"))
}
