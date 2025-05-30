package main

import (
	"errors"
	"fmt"
	"os"
)

type argy struct {
	index int
	value string
}

func newArgy(i int, v string) argy {
	return argy{i, v}

}

func getArgies(args []string) ([]argy, error) {
	if len(args) > 1 {
		argies := make([]argy, len(args)-1)
		for i := 1; i < len(args); i++ {
			argies[i-1] = newArgy(i, args[i])
		}
		return argies, nil
	}
	return nil, errors.New("Need arguments..")
}

func main36() {
	// aNames := [5]string{"Paul", "Joe", "Jan", "Mark", "Sid"}
	// mapNames := arrayToMap(aNames)
	// fmt.Println(mapNames)

	// amap := make(map[int]string, 4)
	// startChar := "!"
	// for i := 0; i < 4; i++ {
	// 	amap[i] = "Paul" + string(startChar[0]+byte(i+23)) // strconv.QuoteRuneToASCII(rune(i+54))
	// }
	// mapToSlices(amap)

	a, e := getArgies(os.Args)
	if e != nil {
		fmt.Println("error: ", e.Error())
	}
	for _, v := range a {
		fmt.Printf("%v, ", v)
	}
}

func mapToSlices(mapi map[int]string) {
	x := len(mapi)

	akeys := make([]int, x)
	avals := make([]string, x)
	c := 0
	for k, v := range mapi {
		akeys[c] = k
		avals[c] = v
		c++
	}
	for _, v := range akeys {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	for _, v := range avals {
		fmt.Printf("%s ", v)
	}

}

func arrayToMap(arr [5]string) map[string]string {
	res := make(map[string]string, len(arr))
	for _, v := range arr {
		res[v] = v
	}
	return res
}
