package termboxUtil

import (
	"github.com/nsf/termbox-go"
)

// ASCIIArt is a []string with more functions
type ASCIIArt struct {
	contents []string
	x, y     int
	bg, fg   termbox.Attribute
}

// CreateASCIIArt Create an ASCII art object from a string slice
func CreateASCIIArt(c []string, x, y int, fg, bg termbox.Attribute) *ASCIIArt {
	i := ASCIIArt{contents: c, x: x, y: y, fg: fg, bg: bg}
	return &i
}

// GetX Return the x position of the modal
func (i *ASCIIArt) GetX() int { return i.x }

// SetX set the x position of the modal to x
func (i *ASCIIArt) SetX(x int) *ASCIIArt {
	i.x = x
	return i
}

// GetY Return the y position of the modal
func (i *ASCIIArt) GetY() int { return i.y }

// SetY Set the y position of the modal to y
func (i *ASCIIArt) SetY(y int) *ASCIIArt {
	i.y = y
	return i
}

// GetBackground Return the current background color of the modal
func (i *ASCIIArt) GetBackground() termbox.Attribute { return i.bg }

// SetBackground Set the current background color to bg
func (i *ASCIIArt) SetBackground(bg termbox.Attribute) *ASCIIArt {
	i.bg = bg
	return i
}

// GetForeground Return the current foreground color
func (i *ASCIIArt) GetForeground() termbox.Attribute { return i.fg }

// SetForeground Set the foreground color to fg
func (i *ASCIIArt) SetForeground(fg termbox.Attribute) *ASCIIArt {
	i.fg = fg
	return i
}

// Align Align the Ascii art over width width with alignment a
func (i *ASCIIArt) Align(a TextAlignment, width int) *ASCIIArt {
	// First get the width of the longest string in the slice
	var newContents []string
	incomingLength := 0
	for _, line := range i.contents {
		if len(line) > incomingLength {
			incomingLength = len(line)
		}
	}
	for _, line := range i.contents {
		newContents = append(newContents, AlignText(AlignText(line, incomingLength, AlignLeft), width, a))
	}
	i.contents = newContents
	return i
}

// HandleKeyPress accepts the termbox event and returns whether it was consumed
func (i *ASCIIArt) HandleKeyPress(event termbox.Event) bool {
	return false
}

// Draw outputs the input field on the screen
func (i *ASCIIArt) Draw() {
	drawX, drawY := i.x, i.y
	for _, line := range i.contents {
		DrawStringAtPoint(line, drawX, drawY, i.fg, i.bg)
		drawY++
	}
}
