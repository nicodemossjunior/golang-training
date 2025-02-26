package main

import "fmt"

func fn01()  {
	fmt.Println("Running function 01")
}

func fn02()  {
	fmt.Println("Running function 02")
}

func main()  {
	defer fn01()
	fn02()
}