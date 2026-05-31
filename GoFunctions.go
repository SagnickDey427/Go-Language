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

//Function with defined return type
func divide(a, b int) (result int, err error) { //Here we have defined the return types as named return values. We can also do this without naming the return values and just specifying the types.
	if b == 0 {
		err = fmt.Errorf("division by zero is not allowed")
		return
	}
	result = a / b
	return
}

//Spread operator in function parameters
func sum(nums ...int) int { //Here nums is a slice of integers that can take any number of integer arguments. We can also use this with other datatypes like string, float, etc.
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
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

	//One liner function using anonymous function
	result3 := applyOperation(4, 6, func(x, y int) int { return x * y })
	fmt.Println(result3)
	//Alternatively , we can also do this
	//  func1:= func(x,y int) int{ return x*y}
	//  result3:= applyOperation(4, 6, func1)
	//  fmt.Println(result3)

	//Example of using a function with defined return type
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Printf("%T\n", err)
	} else {
		fmt.Println("Result:", res)
	}

	//Example of using a function with spread operator in parameters
	//total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	total := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}...) //Here we are using the spread operator to pass a slice of integers as individual arguments to the sum function. This can be done only when the function parameter is defined with the spread operator
	fmt.Println("Total:", total)
}
