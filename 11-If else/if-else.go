package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("less than or equal to 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("Number is greater than zero")
	} else if otherNumber < -10 {
		fmt.Println("Number is less than -10")
	} else {
		fmt.Println("Number between 0 and -10")
	}


}
