package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	street string
	number uint8
}

func main()  {
	fmt.Println("Arquivo structs")

	var u user
	u.name = "David"
	u.age = 21
	fmt.Println(u)

	addressExample := address{"5 avenue", 100}

	u2 := user{"Lucas", 22, addressExample}
	fmt.Println(u2)

	u3 := user{age: 20}
	fmt.Println(u3)
}