package main

import "fmt"

func main()  {
	var var1 string = "Variable 1"
	var2 := "Variable 2"
	fmt.Println(var1)
	fmt.Println(var2)
	
	var (
		var3 string = "hahaha"
		var4 string = "hehehe"
	)

	fmt.Println(var3)
	fmt.Println(var4)

	var5, var6 := "Variable 5", "Variable 6"

	fmt.Println(var5, var6)

	const constant1 string = "Constant 1"
	fmt.Println(constant1)

	var5, var6 = var6, var5
	fmt.Println(var5, var6)
}