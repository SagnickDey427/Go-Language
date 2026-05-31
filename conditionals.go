package main

import "fmt"

func tryCondition() {
	age := 20
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 {
		fmt.Println("You are an adult.")
	}

	//Switch case -- we can use conditionals inside
	num := 12
	switch {
	case num%2 == 0:
		fmt.Println("The number is even.") //Don't need break , breaks automatically after each case in Go
	default:
		fmt.Println("The number is odd.")
	}

	//Switch case -- using fallthrough
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
		fallthrough
	case "Tuesday":
		fmt.Println("It's Tuesday.")
		fallthrough
	case "Wednesday":
		fmt.Println("It's Wednesday.")
	}

	//Switch case -- using multiple values to match a single case
	fruit := "apple"
	switch fruit {
	case "apple", "banana", "orange":
		fmt.Println("It's a common fruit.")
	default:
		fmt.Println("It's an uncommon fruit.")
	}
}
