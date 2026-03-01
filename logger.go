package main

import "log"

func logInfo(message string) {
	log.Printf("[INFO]: %s\n", message) 
}

func logErr(message string) {
	log.Printf("[ERROR]: %s\n", message)
}

func logFatal(message string) {
	log.Printf("[FATAL]: %s\n", message) 
}
