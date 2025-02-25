package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Increasing i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Increasing j", j)
	// 	time.Sleep(time.Second)
	// }

	// names := [3]string{"John", "David", "Lucas"}
	// for index, name := range names {
	// 	fmt.Println(index, name)
	// }

	// for _, name := range names {
	// 	fmt.Println(name)
	// }

	for index, character := range "WORDS" {
		fmt.Println(index, string(character))
	}

	user := map[string]string{
		"name": "Leonardo",
		"lastname": "Silva",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}

}
