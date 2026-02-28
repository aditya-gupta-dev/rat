package main

import (
	"os"

	"golang.design/x/clipboard"
)

func WriteToClipboard(filename string) error {
	err := clipboard.Init()
	if err != nil {
		return err 
	}

	data := clipboard.Read(clipboard.FmtText)

	return os.WriteFile(filename, data, 0o644)
}

func AppendToClipboard(filename string) error {
	err := clipboard.Init()
	if err != nil {
		return err 
	}

	data := clipboard.Read(clipboard.FmtText)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)

	if err != nil { 
		return err 
	}
	defer file.Close()

	_, err = file.Write(data)

	return err 
}
