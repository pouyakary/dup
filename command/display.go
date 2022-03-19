package command

import "dup/terminal"

func DisplayResults(results []string, term *terminal.Terminal) {
	for _, file := range results {
		term.PrintLine(file)
	}
}
