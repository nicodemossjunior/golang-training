package main

import "fmt"

func mathCalcs(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	sum, subtraction := mathCalcs(10, 20)
	fmt.Println(sum, subtraction)
}
