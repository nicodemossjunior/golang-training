package main

import "fmt"

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Widnesday"
	case 5:
		return "Thursday"
	case 6:
		return "Fryday"
	case 7:
		return "Saturday"
	default:
		return "Invalid Number"

	}
}

func dayOfWeek2(number int) string {
	var day string

	switch {
	case number == 1:
		day = "Sunday"
		fallthrough // go to the next condition
	case number == 2:
		day = "Monday"
	case number == 3:
		day = "Tuesday"
	case number == 4:
		day = "Widnesday"
	case number == 5:
		day = "Thursday"
	case number == 6:
		day = "Fryday"
	case number == 7:
		day = "Saturday"
	default:
		day = "Invalid Number"
	}
	return day
}

func main() {
	fmt.Println("Switch")

	day := dayOfWeek(1)
	fmt.Println(day)

	day2 := dayOfWeek2(1)
	fmt.Println(day2)
}
