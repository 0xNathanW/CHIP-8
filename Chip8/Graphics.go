package CHIP8

import (
	"github.com/nsf/termbox-go"
)

const (
	//  Foreground and background colours
	bg = termbox.ColorDefault
	fg = termbox.ColorBlue
)

type Display struct {
	fg, bg termbox.Attribute
}

func InitDisplay() *Display {
	return &Display{fg, bg}
}

// Draw pixels to terminal display
func (d *Display) Draw(c *Chip8) {
	// Loop through pixels
	for y := 0; y < len(c.PixelArray); y++ {
		for x := 0; x < len(c.PixelArray[y]); x++ {
			//  Set colour if pixel is on
			if c.PixelArray[y][x] == 0 {
				termbox.SetCell(x*2, y, ' ', fg, bg)
				termbox.SetCell(x*2+1, y, ' ', fg, bg)
			} else {
				termbox.SetCell(x*2, y, ' ', fg, fg)
				termbox.SetCell(x*2+1, y, ' ', fg, fg)
			}
		}
	}
	// Update terminal
	termbox.Flush()
}
