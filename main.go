package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

func main() {
	var args ArgsModel

	arg.MustParse(&args)

	var size = len(args.Data)

	if args.Clipboard { 
		WriteToClipboard(&args)
		os.Exit(1)
	}

	for _, filename := range args.Filenames {
		err := os.WriteFile(filename, []byte(args.Data), 0644)
		if err != nil { 
			fmt.Fprintf(os.Stderr, "error: failed to write to %s > %s\n",filename, err)
		} else {
			fmt.Fprintf(os.Stdout, "info: written %d bytes to %s\n", size, filename) 
		} 
	}
}
