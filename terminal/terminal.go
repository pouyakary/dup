package terminal

import (
	"fmt"
)

type Terminal struct {
	display bool
}

func NewTerminal(display bool) *Terminal {
	return &Terminal{
		display: display,
	}
}

func (terminal *Terminal) Print(message string) {
	if terminal.display {
		fmt.Print(message)
	}
}

func (terminal *Terminal) PrintLine(message string) {
	if terminal.display {
		fmt.Println(message)
	}
}

func (terminal *Terminal) MoveToX(x int) {
	fmt.Printf("\x1b[%dG", x)
}
