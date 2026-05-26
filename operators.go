package main

import (
	"fmt"
	"math"
	"strconv"
)

func operatorFunc() {
	// x := 2
	// y := 4
	// println(x + y)  //Addition
	// println(x - y)  //Subtraction
	// println(x * y)  //Multiplication
	// println(x / y)  //Division
	// println(x % y)  //Modulus
	// println(x & y)  //Bitwise AND
	// println(x | y)  //Bitwise OR
	// println(x ^ y)  //Bitwise XOR
	// println(x << 1) //Left shift
	// println(x >> 1) //Right shift

	//z:= uint8(7) + int(1000) //Type conversion is required when performing operations on different types. In this case, we convert 1000 to uint8 before adding it to z.

	//z:= 1000 / 7 //This will give us the quotient of the division, which is 142. For actual division we need : z := float32(1000) / float32(7) which will give us 142.85714

	//z := "hi" + string(65) //string() gives ascii value of 65 , so hiA , not hi65

	// z := "hi" + fmt.Sprint(65) //Sprint converts the integer 65 to its string representation "65".
	// fmt.Println(z)

	// x := 5
	// x++
	// x--
	// fmt.Println(x)

	fmt.Println(math.Sqrt(9)) //Using math package

	// w := "1234hello"
	// z, err := strconv.Atoi(w) //Convert string "65" to integer 65
	// fmt.Println(z, err)       //If the conversion is successful, err will be nil. If the conversion fails (e.g., if the string does not represent a valid integer), err will contain an error message.

	x := "1234"
	z, err := strconv.ParseInt(x, 10, 0) //(string , base, bitsize)
	fmt.Println(z, err)
}
