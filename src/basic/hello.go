package main

import "fmt"
import "mymath"

const (
	a = iota
	b = iota
)

func main() {
	fmt.Printf("Hello, world. Sqrt(2) = %v\n\n", mymath.Sqrt(2))
	fmt.Printf("a=%d, b=%d", a, b)
}
