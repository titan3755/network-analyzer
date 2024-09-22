package utils

import (
	"github.com/inancgumus/screen"
)

// ResetTerminal is a function that clears the terminal screen and moves the cursor to the top left corner

func ResetTerminal() {
	screen.Clear()
	screen.MoveTopLeft()
}