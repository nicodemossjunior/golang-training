package main

import "fmt"

type pearson struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	pearson
	course  string
	college string
}

func main() {
	fmt.Println("Inheritance")

	p1 := pearson{"John", "Petrus", 20, 180}
	fmt.Println(p1)

	e1 := student{p1, "Engineering", "USP"}
	fmt.Println(e1)

	fmt.Println(e1.name, e1.age)

}
