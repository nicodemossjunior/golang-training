package main

import (
	"fmt"
)

func main() {
	user := map[string]string{
		"name":     "John",
		"lastname": "Mayer",
	}

	fmt.Println(user)

	user2 := map[string]map[string]string{
		"name": {
			"first": "Petrus",
			"last":  "Willian",
		},
		"course": {
			"name":   "Engineering",
			"campus": "campus 1",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)
	
	user2["ticker"] = map[string]string{
		"name": "PETR4",
		"fullname": "Petrobras",
	}
	
	fmt.Println(user2)
}
