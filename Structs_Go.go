package main

import "fmt"

type Person struct { //This 'Person' will be treated as a new dataType and can be used like any other Go datatype.
	name   string
	age    uint
	gender string
	salary float64
}

func tryStructs() {
	//Empty struct
	p1 := Person{} //We can access p1.name , p1.age etc. it's fields
	fmt.Println(p1)

	//Struct with values -- labelled input
	p2 := Person{name: "Sagnick", age: 20, gender: "Male", salary: 1000000.0}
	fmt.Println(p2)

	//Struct with values -- unlabelled input
	p3 := Person{"Sahil", 21, "Male", 50000.0}
	fmt.Println(p3)
}
