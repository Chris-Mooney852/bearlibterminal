package main

import (
	"C"
	blt "bearlibterminal"
	"strconv"
)

const (
	WINDOW_SIZE_X = 100
	WINDOW_SIZE_Y = 35
	TITLE         = "BearRogue"
	FONT          = "fonts/UbuntuMono.ttf"
	FONT_SIZE     = 24
)

func init() {
	blt.Open()

	// Setup window config settings
	size := "size=" + strconv.Itoa(WINDOW_SIZE_X) + "x" + strconv.Itoa(WINDOW_SIZE_Y)
	title := "title='" + TITLE + "'"
	window := "window: " + size + "," + title

	// Font Config
	fontSize := "size=" + strconv.Itoa(FONT_SIZE)
	font := "font: " + FONT + ", " + fontSize

	blt.Set(window + "; " + font)
	blt.Clear()
}

func main() {
	blt.Print(1, 1, "Hello, world!")
	blt.Refresh()

	for blt.Read() != blt.TK_CLOSE {
		// For now, do nothig
	}

	// If the user hits the clone button, close the window and exit
	blt.Close()
}
