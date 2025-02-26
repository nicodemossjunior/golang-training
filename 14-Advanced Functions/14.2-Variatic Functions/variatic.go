package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func write(text string, nums ...int)  {
	for _, num := range nums {
		fmt.Println(text, num)
	}
}

func main() {
	totalSum := sum(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalSum)

	write("Hello!!!", 10, 2, 3, 4, 5, 6, 7)
}
