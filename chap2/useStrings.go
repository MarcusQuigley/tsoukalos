package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

type Power2 int

func main23() {
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))
	f("%s\n", s.Split("abcd efgh", ""))
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	f("%s\n", s.Replace("abcd efgh", "", ",", 2))
	f("%s\n", s.Replace("abcd efgh", "", ",", -1))

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		p2_5
		p2_6
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
	fmt.Println("p2_5:", p2_5)
}
