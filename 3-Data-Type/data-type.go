package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEGER NUMBERS
	var number int64 = -1000000000000000000
	fmt.Println(number)

	var number2 uint32 = 100000000
	fmt.Println(number2)

	// alias
	// INT32 = RUNE
	var number3 rune = 123456
	fmt.Println(number3)

	// BYTE = UINT8
	var number4 byte = 123
	fmt.Println(number4)
	// END INTEGER NUMBERS

	// REAL NUMBERS
	var realNumber1 float32 = 123.45
	fmt.Println(realNumber1)

	var realNumber2 float64 = 1230004540500.45
	fmt.Println(realNumber2)

	realNumber3 := 1230500.45
	fmt.Println(realNumber3)
	// END REAL NUMBERS

	// STRINGS
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	// RETURNS VALUE FROM ASCII TABLE
	char := 'A'
	fmt.Println(char)

	// END STRINGS

	text := 5
	fmt.Println(text)

	var boolean bool
	fmt.Println(boolean)

	var err error = errors.New("Internal error")
	fmt.Println(err)

}
