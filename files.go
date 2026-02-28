package main

import "os"

func WriteToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0o644)
}

func AppendToFile(filename string, data []byte) error { 
	file, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil { 
		return err 
	}
	defer file.Close() 

	_, err = file.Write(data)

	return err 
}
