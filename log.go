package nogo

import (
	"log"
)

func LogPrintln(v ...interface{}) {
	log.Println(v)
}

func LogFatalln(v ...interface{}) {
	log.Fatalln(v)
}
