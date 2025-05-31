package main

import (
	"errors"
	"fmt"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

type TreeLast[T any] []T

func (t TreeLast[T]) replaceLast(el T) (TreeLast[T], error) {
	if len(t) == 0 {
		return t, errors.New("its empty")
	}
	t[len(t)-1] = el
	return t, nil
}

func PrintSlice[B any](s []B) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

func Same[D comparable](x, y D) bool {
	return x == y
}

func Add[T Numeric](x, y T) T {
	return x + y
}

func main4() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"ws", "ed", "gfd", "bg"})
	fmt.Println("1==2? ", Same(1, 2))
	fmt.Println("2==2? ", Same(2, 2))
	fmt.Println("a=a? ", Same("a", "a"))
	fmt.Println(Add(4, 6))
	fmt.Println(Add(4.0, 6.3))

	ts := TreeLast[string]{"aa", "bb"}
	fmt.Println(ts)
	ts1, e := ts.replaceLast("cc")
	if e != nil {
		fmt.Println("err: ", e.Error())
	}
	fmt.Println(ts1)
}
