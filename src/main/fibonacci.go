package main

import "fmt"

func main()  {
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(50))
}

func fibonacci(n int) (*[]int) {
	result := make([]int, n)
	if n >= 1 {
		result[0] = 1
	}

	if n >= 2 {
		result[1] = 1
	}

	if n <= 2 {
		return &result
	}

	for i := 2; i < n; i++ {
		result[i] = result[i - 1] + result[i - 2]
	}

	return &result
}
