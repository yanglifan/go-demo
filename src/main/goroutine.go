package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(2 * time.Second)

	ci := make(chan int)
	go readFromChannel(ci)
	go writeIntoChannel(ci)
	time.Sleep(2 * time.Second)
}

func readFromChannel(ci chan int) {
	i := <-ci
	fmt.Println("Read from channel ->", i)
}

func writeIntoChannel(ci chan int) {
	ci <- 1
}