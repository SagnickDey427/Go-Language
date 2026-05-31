package main

import "fmt"

//A basic function in go looks like this
func add(a int, b int) int {
	return a + b
}

//Function returning multiple values
func swap(x, y string) (string, string) { //x,y string is a shorthand for x string, y string
	return y, x
}

//Higher order function -- taking a func as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

//Higher order function -- function returning another function
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func tryFuncyions() {
	val := add(2, 3)
	fmt.Println(val)
	fmt.Printf("%T\n", add) // type of add is func(int, int) int

	// Example of using a function that returns multiple values
	str1, str2 := swap("Hello", "World")
	fmt.Println(str1, str2)

	// Example of using a higher order function
	result := applyOperation(5, 10, add)
	fmt.Println(result)

	// Example of using a higher order function that returns another function
	mult := multiplier(3)
	result2 := mult(7)
	fmt.Println(result2)
}
