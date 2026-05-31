package main

import "fmt"

func tryMaps() {
	//var mp map[string]int = map[string]int{"one":1,"two":2}

	mp := map[string]int{"one": 1, "two": 2}
	fmt.Println(mp)

	//Using make
	mp2 := make(map[string]int)
	mp2["three"] = 3 //Adding a new key to the map
	mp2["four"] = 4
	fmt.Println(mp2)

	mp3 := make(map[string][]int) //string : slice
	mp3["numbers"] = []int{1, 2, 3, 4, 5}
	fmt.Println(mp3)
	mp3["numbers"] = []int{} //Updating the value of an existing key
	fmt.Println(mp3)

	//Deleting a key from the map
	delete(mp, "one") //In-place deletion so unlike append, it returns nothing
	fmt.Println(mp)

	//Checking if a key exists in the map
	value, exists := mp["two"] //exists = true if exists , flase otherwise and value  = default value of the datatype if key does not exist
	fmt.Println(value, exists)
	value, exists = mp["one"]
	fmt.Println(value, exists)

}
