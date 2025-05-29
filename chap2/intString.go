package main

import (
	"fmt"
	"os"
	"strconv"
)

func main22() {
	if len(os.Args) == 1 {
		fmt.Println("enter an int")
		return
	}
	n, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println("enter a valid int")
		return
	}

	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %v of type %T\n", input, input)
	input = string(n)
	fmt.Printf("string() %v of type %T\n", input, input)

}
