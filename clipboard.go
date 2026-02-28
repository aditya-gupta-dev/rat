package main

import (
	"log"
	"os"

	"golang.design/x/clipboard"
)

func WriteToClipboard(args *ArgsModel) { 

	err := clipboard.Init()	

	if err != nil { 
		log.Fatalf("error: %s\n", err)
	}

	data := clipboard.Read(clipboard.FmtText) 

	for _, filename := range args.Filenames {
		os.WriteFile(filename, data, 0644)
	}
}
