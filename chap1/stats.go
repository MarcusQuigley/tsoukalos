package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 0 {
		fmt.Println("need args")
		return
	}
	var min, max, sum float64
	var initialized = 0
	nValues := 0
	for i := 1; i < len(args); i++ {
		n, e := strconv.ParseFloat(args[i], 64)
		if e != nil {
			continue
		}
		nValues = nValues + 1
		sum = sum + n
		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("number of values:", nValues)
	fmt.Println("min", min)
	fmt.Println("max", max)

	if nValues == 0 {
		return
	}

	meanValue := sum / float64(nValues)
	fmt.Printf("mean: %.5f\n", meanValue)

	var squared float64
	for i := 1; i < len(args); i++ {
		n, e := strconv.ParseFloat(args[i], 64)
		if e != nil {
			continue
		}
		squared = squared + math.Pow((n-meanValue), 2)

	}
	sd := math.Sqrt(squared / float64(nValues))
	fmt.Printf("standard deviation %.5f\n", sd)
}
