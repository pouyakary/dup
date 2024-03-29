package command

import (
	"fmt"
	"os"
)

type ComputationContext struct {
	Directory   string
	DisplayHelp bool
	Quiet       bool
	Remove      bool
	Exact       bool
}

func ParseCommandLineArguments() *ComputationContext {
	var (
		directory        string
		help             = false
		quiet            = false
		remove           = false
		exact            = false
		args             = os.Args[1:]
		foundDirectories = 0
	)

	for _, arg := range args {
		switch arg {
		case "-h", "--help", "help":
			help = true
		case "-q", "--quiet", "quiet":
			quiet = true
		case "-r", "--remove", "remove":
			remove = true
		case "-e", "exact":
			exact = true
		default:
			directory = arg
			foundDirectories++
		}
	}

	if help {
		return &ComputationContext{
			Directory:   "",
			DisplayHelp: true,
			Quiet:       false,
			Remove:      false,
			Exact:       true,
		}
	}

	if foundDirectories == 0 {
		fmt.Println("no directory specified")
		os.Exit(1)
	}
	if foundDirectories > 1 {
		fmt.Println("more than one directory specified")
		os.Exit(1)
	}

	dir, err := os.Stat(directory)
	if os.IsNotExist(err) {
		fmt.Println("path does not exist")
		os.Exit(1)
	}
	if !dir.IsDir() {
		fmt.Println("path is not a directory")
		os.Exit(1)
	}

	return &ComputationContext{
		Directory:   directory,
		DisplayHelp: help,
		Quiet:       quiet,
		Remove:      remove,
		Exact:       exact,
	}
}
