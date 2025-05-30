package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrDivideByZero = errors.New("divide by zero error")

func main33() {

	// sl := arraysToSlice()
	// for _, v := range sl {
	// 	fmt.Printf("%d ", v)
	// }
	//arraysToArray()

	if len(os.Args) == 1 {
		fmt.Println("Please enter a  number.")
		return
	}
	strNum := os.Args[1]
	num, e := strconv.Atoi(strNum) //.ParseInt(strNum, 10, 64)
	if e != nil {
		fmt.Printf("error with num args %s. :%v", strNum, e.Error())
		return
	}
	checkerrorIs(num)
}

func arraysToSlice() []int {
	a1 := [4]int{1, 2, 3, 4}
	var a2 = [4]int{5, 6, 7, 8}

	var result []int
	result = append(result, a1[:]...)
	result = append(result, a2[:]...)

	return result
}

func arraysToArray() {
	a1 := [4]int{1, 2, 3, 4}
	var a2 = [4]int{5, 6, 7, 8}
	var ar [8]int
	for i := 0; i < 4; i++ {
		ar[i] = a1[i]
		ar[i+4] = a2[i]
	}
	for _, v := range ar {
		fmt.Printf("%d ", v)
	}

}

func divide(x, y int) (int, error) {
	if x == 0 || y == 0 {
		return 0, ErrDivideByZero
	}
	return (x / y), nil
}

func checkerrorIs(y int) {
	x := 4
	result, e := divide(x, y)
	if e != nil {
		if errors.Is(e, ErrDivideByZero) {
			fmt.Println("dont use zero u twat")
		} else {
			fmt.Println("error: ", e.Error())
		}

		return
	}
	fmt.Println("result: ", result)

}
