package main

import (
	"log"
	//"log/syslog"
	"os"
)

func main3() {
	// syslog, e := syslog.New(syslog.LOG_SYSLOG, "syslog.go")
	// if e != nil {
	// 	log.Println(e)
	// 	return
	// } else {
	// 	log.SetOutput(syslog)
	// 	log.Print("Everythings fine!")
	// }

	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}
