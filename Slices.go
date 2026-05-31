package main

import "fmt"

func trySlices() {
	//Pointer --> Starting index of the slice in the underlying array
	//Length --> Number of elements in the slice
	//Capacity --> Maximum number of elements that the slice can hold

	//Slices are dynamic arrays that can grow and shrink in size. They are built on top of arrays and provide more flexibility.
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] //This creates a slice that includes elements from index 1 to 3 (4 is exclusive) of the array. So slice1 will contain [2, 3, 4]. Just like python slicing

	fmt.Println(slice1, len(slice1), cap(slice1)) //Length of slice is 3 and capacity is 4 because it can grow upto the end of the array

	slice2 := slice1[:4] //Since cap(slice1) >= 4, we can extend the slice to include more elements from the underlying array. So slice2 will contain [2, 3, 4, 5]

	slice2[0] = 100 //Slice2 is slice of the same array (not slice1) so changing slice2 will change original array and since original array changed , so slice1 also changes.
	fmt.Println(slice2, slice1, arr)

	slice3 := []string{"Go", "is", "awesome"} //Behind the scenes , Go creates an array of 3 strings and slice3 is a slice that points to that array.

	for i := 0; i < 3; i++ {
		slice3 = append(slice3, fmt.Sprintf("new%d", i)) //append adds new elements to the end of the slice. If the underlying array is not large enough to hold the new elements, a new array is allocated with double capacity (just like vectors in C++) and the existing elements are copied over.
		fmt.Println(slice3, len(slice3), cap(slice3))
	}

	//Using spread operator in append()
	slice4 := append([]byte("hello "), "world"...) //The "..." is the spread operator that allows us to append each character of the string "world" as individual bytes to the slice. So slice4 will contain the bytes of "hello world". We can also use it to append another slice, for example: slice5 := append(slice4, slice3...) will append all elements of slice3 to slice4.

	fmt.Println(string(slice4)) //Since slice4 is a slice of bytes, we need to convert it back to string before printing.

	slice5 := make([]string, 3, 5) //make creates a slice with specified length and capacity.

	slice5[0] = "Go"
	slice5[1] = "is"
	slice5[2] = "fun"

	fmt.Println(slice5)

	test(slice5) //Since slices are reference types, when we pass slice5 to the test function, we are passing a reference to the underlying array. So any changes made to the slice inside the test function will affect the original slice outside the function.

	fmt.Println(slice5)
}

func test(sl []string) {
	sl[0] = "changed"
}
