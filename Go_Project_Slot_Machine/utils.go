package main

import (
	"fmt"
)

func GetName() string {
	fmt.Println("🌟 Welcome to Dey's Casino....")
	fmt.Printf("👉 Enter your name: ")
	fmt.Print("> ")
	name := ""
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("🪸 Error in taking name input")
		return ""
	}
	fmt.Println("Welcome to Dey's Casino Mr.", name, ", let's have some fun! 🎰")
	return name
}
