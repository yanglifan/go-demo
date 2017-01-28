package main

import (
	"fmt"
)

type NameAge struct {
	Name string
	age  int
}

func main() {
	a := &NameAge{Name: "yanglifan", age:30}
	fmt.Print(a.Name)
}