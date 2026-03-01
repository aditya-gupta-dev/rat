package main

import "log"

func logInfo(message string) {
	log.Printf("info: %s\n", message) 
}

func logErr(message string) {
	log.Printf("error: %s\n", message)
}

func logFatal(message string) {
	log.Printf("fatal: %s\n", message) 
}
