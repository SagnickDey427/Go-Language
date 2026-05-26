package main

import "fmt"

func formatFunc() {
	x := 2
	fmt.Println(345, "Output", x)
	fmt.Println("This is line 2 of file formats_fmt.go") //Println gives automatic '\n' character at the end.

	fmt.Printf("The value of x is : %v and it's type is : %T and it's binary value is : %b\n", x, x, x) //Format strings in Go act just like C ;  %T = Type of var, %v = Value , %b = binary form of value , %s for string ,%% for %, %f for float and so on. You can also use width and precision with these format verbs. For example, %.2f will format a float to two decimal places, %10.2f will format a float to two decimal places and ensure it takes up at least 10 characters in width, padding with spaces if necessary.
	a := 3.14567
	y := fmt.Sprintf("%10.2f\"", a) //Sprintf is used to format a string and return it instead of printing it. It takes the same format verbs as Printf.
	fmt.Println(y, y)
}
