package main

import "fmt"

func mainX() {
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Your name is %v\n", name)
}
