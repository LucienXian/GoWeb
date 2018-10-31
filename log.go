package tinyWeb

import (
	"log"
	"os"
	"io"
	"io/ioutil"
)

var (
	Trace	*log.Logger //All logs
	Info	*log.Logger 	//Important info
	Warning	*log.Logger//Pay attention
	Error	*log.Logger	//Error
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file: ", err)
	}
	
	Trace = log.New(ioutil.Discard, 
		"TRACE: ", 
		log.Ldate | log.Ltime | log.Lshortfile)
	
	Info = log.New(os.Stdout, 
		"INFO: ", 
		log.Ldate | log.Ltime | log.Lshortfile)
	
	Warning = log.New(os.Stdout, 
		"WARNING: ", 
		log.Ldate | log.Ltime | log.Lshortfile)
	
	Error = log.New(io.MultiWriter(file, os.Stderr), 
		"ERROR: ", 
		log.Ldate | log.Ltime | log.Lshortfile)
}