package main

import "fmt"

type I interface {
	Set(int)
	Get() int
}

func demoInterfaceParam(i I) {
	i.Set(1)
	fmt.Println(i.Get())
}

func demoEmptyInterface(something interface{}) int {
	return something.(I).Get()
}