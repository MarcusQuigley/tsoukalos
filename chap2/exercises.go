package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

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
	errorIs(num)
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

func errorIs(y int) {
	x := 4
	result := x / y
	fmt.Println("result: ", result)

}
