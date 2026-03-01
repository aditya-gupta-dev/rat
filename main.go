package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

func main() {
	var args ArgsModel

	arg.MustParse(&args)

	data := []byte(args.Data)
	size := len(data)

	if args.Clipboard && size != 0 {

		cmdArgs := os.Args[1:]
		clipArgPos := LinearSearch(cmdArgs, "-c")

		if clipArgPos < LinearSearch(cmdArgs, "-d") || clipArgPos < LinearSearch(cmdArgs, "--data") {
			for _, filename := range args.Filenames {
				if err := AppendToClipboard(filename); err != nil {
					logErr(fmt.Sprintf("failed to append clipboard to %s, %s", filename, err))
				}
				if err := AppendToFile(filename, data); err != nil {
					logErr(fmt.Sprintf("faield to append to %s %s", filename, err))
				}
			}
		} else {
			for _, filename := range args.Filenames {
				if err := AppendToFile(filename, data); err != nil {
					logErr(fmt.Sprintf("failed to write clipboard to %s %s", filename, err))
				}

				if err := AppendToClipboard(filename); err != nil {
					logErr(fmt.Sprintf("failed to write to %s %s", filename, err))
				}
			}
		}
	} else {
		if args.Clipboard {
			for _, filename := range args.Filenames {
				if err := WriteToClipboard(filename); err != nil {
					logErr(fmt.Sprintf("failed to copy to clipboard %s %s", filename, err))
				}
			}
		} else {
			for _, filename := range args.Filenames {
				if err := WriteToFile(filename, data); err != nil {
					logErr(fmt.Sprintf("failed to write to file %s %s", filename, err))
				}
			}
		}
	}
}
