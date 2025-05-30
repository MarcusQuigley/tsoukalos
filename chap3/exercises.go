package main

import (
	"fmt"
)

func mains() {
	// aNames := [5]string{"Paul", "Joe", "Jan", "Mark", "Sid"}
	// mapNames := arrayToMap(aNames)
	// fmt.Println(mapNames)

	amap := make(map[int]string, 4)
	startChar := "!"
	for i := 0; i < 4; i++ {
		amap[i] = "Paul" + string(startChar[0]+byte(i+23)) // strconv.QuoteRuneToASCII(rune(i+54))
	}
	mapToSlices(amap)
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
