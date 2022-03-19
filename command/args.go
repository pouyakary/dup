package command

import (
	"fmt"
	"os"
)

type ComputationContext struct {
	Directory   string
	DisplayHelp bool
	Quite       bool
	Remove      bool
}

func ParseCommandLineArguments() *ComputationContext {
	var (
		directory        string
		help             = false
		quite            = false
		remove           = false
		args             = os.Args[1:]
		foundDirectories = 0
	)

	for _, arg := range args {
		switch arg {
		case "-h", "--help", "help":
			help = true
		case "-q", "--quite", "quite":
			quite = true
		case "-r", "--remove", "remove":
			remove = true
		default:
			directory = arg
			foundDirectories++
		}
	}

	if help {
		return &ComputationContext{
			Directory:   "",
			DisplayHelp: true,
			Quite:       false,
			Remove:      false,
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
		Quite:       quite,
		Remove:      remove,
	}
}
