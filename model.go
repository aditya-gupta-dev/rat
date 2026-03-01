package main

type ArgsModel struct { 
	Filenames []string `arg:"positional" help:"file names you wanna write to"`
	Data string `arg:"-d" help:"data to be written"`
	Clipboard bool `arg:"-c" help:"write data from clipboard"`
	Verbose bool `arg:"-v" help:"verbose mode write logs"`
}
