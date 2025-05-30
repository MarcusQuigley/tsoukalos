package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main34() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}
	f, t := 0, 0
	for i := 1; i < len(arguments); i++ {
		s := arguments[i]

		ret := matchInt(s)
		if ret {
			t++
		} else {
			f++
		}
	}
	fmt.Printf("totals: true %d, false %d", t, f)

}
