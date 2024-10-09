// main.go
package main

import (
    "fmt"
    "github.com/yourusername/greetings"
)

func main() {
    message := greetings.Hello("Alice")
    fmt.Println(message)
}
