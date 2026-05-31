package main

import "fmt"

func tryLoop() {
	//For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While loop -- Go doesn't have while loop but we can achieve the same functionality using for loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}

func tryString() {
	str1 := "Hello, World!" //Behind the scenes , Go stores strings as an array of 1 byte characters. So each character is converted into its byte code (UTF-8 encoding) and stored in array.
	fmt.Println(str1[0])    // So indexing string is actually indexing 1-byte of that i-th indexed character. It gives problem when we have special characters that take more than 1 byte of space. For example, in "Go🎯" , G and o take 1 byte each but 🎯 takes 4 bytes so indexing it will give us the byte code of the first byte of that character which is not what we want.
	fmt.Println(str1)

	str2 := "Go🎯" //Modern string uses UTF-8 standard that holds upto 4 bytes for each character , it means some special characters , like emojis 🎯 etc. will take upto 4 bytes of place instead of 1 byte.

	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i]) //This will give us the byte code of the first byte of that character which is not what we want.
	}

	//Best way to iterate over a string is to use range loop
	for _, ch := range str2 { //range returns index,character for each character in the string. We can ignore the index by using _.
		fmt.Printf("%c ", ch) //This will give us the correct character even for special characters that take more than 1 byte of space.
	}
}
