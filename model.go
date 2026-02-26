package main

type ArgsModel struct { 
	Filenames []string `arg:"positional" help:"file names you wanna write to"`
	Data string `help:"data to be written"`
}
