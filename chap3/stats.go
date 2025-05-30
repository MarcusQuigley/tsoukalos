package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func readFile(filepath string) ([]float64, error) {
	_, e := os.Stat(filepath)
	if e != nil {
		return nil, e
	}
	f, e := os.Open(filepath)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	data, e := csv.NewReader(f).ReadAll()
	if e != nil {
		return nil, e
	}
	result := make([]float64, len(data))

	for _, v := range data {
		r, e := strconv.ParseFloat(v[0], 64)
		if e != nil {
			log.Println("Error reading: ", v[0], e.Error())
			continue
		}
		result = append(result, r)
	}
	return result, nil
}

func stdDev(x []float64) (float64, float64) {
	sum := 0.0
	for _, val := range x {
		sum = sum + val
	}
	mean := sum / float64(len(x))
	fmt.Printf("Mean value: %.5f\n", mean)
	// sd
	var squared float64
	for i := 0; i < len(x); i++ {
		squared = squared + math.Pow((x[i]-mean), 2)
	}
	sd := math.Sqrt(squared / float64(len(x)))
	return mean, sd

}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}

	return normalized
}

func randomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func main32() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	fp := arguments[1]
	values, e := readFile(fp)
	if e != nil {
		log.Println("error: ", fp, e)
		os.Exit(1)
	}
	sort.Float64s(values)

	fmt.Println("Number of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	mean, sd := stdDev(values)

	fmt.Printf("Standard deviation: %.5f\n", sd)
	normalized := normalize(values, mean, sd)
	fmt.Println("Normalized:", normalized)

}
