package main

import "fmt"

func main() {
	// ARITHMETIC OPERATORS
	addition := 1 + 2
	subtraction := 2 - 2
	division := 12 / 4
	multiplication := 8 * 5
	remindOfDivision := 10 % 2

	fmt.Println(addition, subtraction, division, multiplication, remindOfDivision)

	var number1 int16 = 10
	var number2 int16 = 25
	addition2 := number1 + number2
	fmt.Println(addition2)

	// END ARITHMETIC OPERATORS

	// ASSIGNMENT
	var variable1 string = "String"
	var variable2 string = "String2"
	fmt.Println(variable1, variable2)
	// END ASSIGNMENT

	// RELATIONAL OPERATORS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// END RELATIONAL OPERATORS

	// LOGICAL OPERATORS
	fmt.Println("--------------------")
	varTrue, varFalse := true, false
	fmt.Println(varTrue && varFalse)
	fmt.Println(varTrue || varFalse)
	fmt.Println(!varTrue)
	fmt.Println(!varFalse)
	// END LOGICAL OPERATORS

	// UNARY OPERATORS
	fmt.Println("--------------------")
	number := 10
	number++
	number += 15
	fmt.Println(number)

	number--
	number -= 20

	number *= 3
	number /= 10
	number %= 3

	fmt.Println(number)
	// END UNARY OPERATORS

}
