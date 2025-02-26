package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return 1
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("Recursive Functions")

	// 1 1 2 3 5 8 13 21 34 55

	position := uint(12)

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}
}
