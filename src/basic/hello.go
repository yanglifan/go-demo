package main

import "fmt"
import "mymath"

const (
	a = iota
	b
)

func main() {
	fmt.Printf("Hello, world. Sqrt(2) = %v\n\n", mymath.Sqrt(2))
	fmt.Printf("a=%d, b=%d\n", a, b)

	s := "hello"
	c := []rune(s) // rune is int32 alias
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	for pos, char := range "aux" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}

	loopArray()
}

func loopArray() {
	array := []string{"a", "b", "c"}
	for i := 0; i < len(array); i++ {
		fmt.Printf("%s", array[i])
	}
}
