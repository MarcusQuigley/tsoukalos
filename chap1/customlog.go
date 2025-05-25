package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main4() {
	LOGFILE := path.Join(os.TempDir(), "mgo.log")
	fmt.Println(LOGFILE)
	f, e := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer f.Close()
	lstdFlags := log.Ldate | log.Ltime | log.Lshortfile
	w := io.MultiWriter(f, os.Stderr)
	ilog := log.New(w, "ilog ", lstdFlags) // log.LstdFlags)
	ilog.Println("Logging it ..")
	ilog.Println("shutting down")

}
