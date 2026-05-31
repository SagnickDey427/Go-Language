package main

import "fmt"

func tryArrays() {
	arr := [5]int{1, 2, 34, 5, 6} //Implicit declaration of array
	//Size of arrays are fixed after created, like in C and C++ arrays
	for i, val := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, val)
	}

	arr2 := [...][2]int{ //Multidimensional array with implicit size
		{1, 2},
		{3, 4},
		{5, 6},
	}
	for i, row := range arr2 {
		for j, val := range row {
			fmt.Printf("Index: [%d][%d], Value: %d\n", i, j, val)
		}
	}
}
