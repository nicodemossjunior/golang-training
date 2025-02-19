package main

import "fmt"

func plus(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalcs(n1, n2 int8) (int8, int8) {
	plus := n1 + n2
	minus := n1 - n2
	return plus, minus
}

func main() {
	sum := plus(12, 15)
	fmt.Println(sum)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Text of the function 1")
	fmt.Println(result)

	resultSum, resultMinus := mathCalcs(10, 15)
	fmt.Println(resultSum, resultMinus)

	resultSum1, _ := mathCalcs(20, 10)
	fmt.Println(resultSum1)

}
