package main

import "fmt"

type node[T any] struct {
	Data T
	Next *node[T]
}

type list[T any] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		Next: nil,
	}
	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.Next == nil {
		l.start.Next = &n
		return
	}
	temp := l.start
	l.start = l.start.Next
	l.add(data)
	l.start = temp
}

func main41() {
	var listy list[int]
	fmt.Println(listy)
	listy.add(12)
	listy.add(9)
	listy.add(3)
	listy.add(1)

	cur := listy.start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		cur = cur.Next
	}
}
