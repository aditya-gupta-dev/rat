package main

import (
	"os"

	"github.com/alexflint/go-arg"
)

func main() {
	var args ArgsModel

	arg.MustParse(&args)

	var data = []byte(args.Data)
	size := len(data) 

	if args.Clipboard && size != 0 {

		cmdArgs := os.Args[1:]
		clipArgPos := LinearSearch(cmdArgs, "-c")

		if clipArgPos < LinearSearch(cmdArgs, "-d") || clipArgPos < LinearSearch(cmdArgs, "--data") {
			for _, filename := range args.Filenames {
				// TODO: handle errors 
				AppendToClipboard(filename)
				AppendToFile(filename, data) 
			}
		} else { 
			for _, filename := range args.Filenames { 
				// TODO: handle errors
				AppendToFile(filename, data)
				AppendToClipboard(filename)
			}
		}
	} else { 
		if args.Clipboard { 
			for _, filename := range args.Filenames {
				// TODO: handle errors 
				WriteToClipboard(filename)
			}
		} else {
			for _, filename := range args.Filenames { 
				// TODO: handle errors 
				WriteToFile(filename, data) 
			}
		}
	}


}
