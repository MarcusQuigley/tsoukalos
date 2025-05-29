package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main25() {
	i := int(1)
	fmt.Println("untyped: ", i*untypedConstant)
	fmt.Println("typed: ", int16(i)*typedConstant)
}
