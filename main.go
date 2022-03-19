package main

import (
	"dup/command"
	"dup/detector"
	"dup/terminal"
)

func main() {
	var (
		context    = command.ParseCommandLineArguments()
		term       = terminal.NewTerminal(!context.Quite)
		duplicates = detector.FindDouplicateFiles(context.Directory, term)
	)

	if context.Remove {
		command.RemoveFiles(context.Directory, duplicates)
	}

	command.DisplayResults(duplicates, term)
}
