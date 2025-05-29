package main

import "fmt"

func main21() {
	stringy := "A stringZåß"
	//stringy2 := "A string"
	bitey := []byte(stringy)
	runey := '€'

	fmt.Printf("string: %v\n", len(stringy))
	fmt.Printf("bytes: %v\n", len(bitey))
	fmt.Printf("rune: %s\n", runey)
	fmt.Printf("rune as string: %c\n", runey)
	fmt.Println("rune As an int32 value:", runey)

	for _, v := range stringy {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
	for _, v := range stringy {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
}
