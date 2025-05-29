package main

import "fmt"

type astruct struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	fmt.Printf("... *x: %v, **x: %v  ...\n", x, *x)
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func jesus(xa *float64) {
	x := &xa
	y := *xa
	yy := *x
	fmt.Println(xa)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(yy)

}

func main24() {
	var f float64 = 12.123

	fmt.Println("memory address of f: ", &f)
	fp := &f
	fmt.Println("memory address of f: ", fp)
	fmt.Println("value of f: ", *fp)
	processPointer(fp)
	fmt.Printf("value of f: %.2f\n ", f)
	x := returnPointer(f)
	fmt.Printf("value of x: %.2f\n", *x)
	xx := bothPointers(fp)
	fmt.Printf("value of xx: %.2f\n", *xx)

	var k *astruct
	fmt.Println(k)
	// Therefore you are allowed to do this:
	if k == nil {
		k = new(astruct)
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}

}
