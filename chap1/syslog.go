package main

import (
	"log"
	"log/syslog"
)

func main() {
	syslog, e := syslog.New(syslog.LOG_SYSLOG, "syslog.go")
	if e != nil {
		log.Println(e)
	} else {
		log.SetOutput(syslog)
		log.Print("Everythings fine!")
	}
}
