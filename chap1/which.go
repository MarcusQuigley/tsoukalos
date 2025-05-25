package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main2() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please enter an argument!")
		return
	}
	fileName := args[1]
	path := os.Getenv("PATH")
	paths := filepath.SplitList(path)

	for _, v := range paths {
		fullPath := filepath.Join(v, fileName)
		fileInfo, e := os.Stat(fullPath)
		if e != nil {
			continue
		}
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			continue
		}
		//is executable
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}

}
