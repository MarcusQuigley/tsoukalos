package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
	_, e := os.Stat(filepath)
	if e != nil {
		return nil, e
	}
	f, e := os.Open(filepath)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '#'
	r.TrimLeadingSpace = true
	// l, e := r.Read()
	// if e != nil {
	// 	return nil, e
	// }
	// fmt.Println(l)
	lines, e := csv.NewReader(f).ReadAll()
	if e != nil {
		return nil, e
	}
	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvFile, e := os.Create(filepath)
	if e != nil {
		return e
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	writer.Comma = '\t'

	for _, r := range myData {
		t := []string{r.Name, r.Surname, r.Number, r.LastAccess}
		e = writer.Write(t)
		if e != nil {
			return e
		}
	}
	writer.Flush()
	return nil
}

func main35() {
	if len(os.Args) != 3 {
		log.Println("csvData input output")
		os.Exit(1)
	}
	input := os.Args[1]
	output := os.Args[2]
	lines, e := readCSVFile(input)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	for _, l := range lines {

		line := strings.Split(l[0], "#")
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		log.Println(temp)
	}
	e = saveCSVFile(output)
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
