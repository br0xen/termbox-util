package termboxUtil

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
)

// TextAlignment is an int value for how we're aligning text
type TextAlignment int

const (
	// AlignLeft Aligns text to the left
	AlignLeft = iota
	// AlignCenter Aligns text to the center
	AlignCenter
	// AlignRight Aligns text to the right
	AlignRight
)

/* Basic Output Helpers */

// DrawStringAtPoint Draw a string of text at x, y with foreground color fg, background color bg
func DrawStringAtPoint(str string, x int, y int, fg termbox.Attribute, bg termbox.Attribute) (int, int) {
	xPos := x
	for _, runeValue := range str {
		termbox.SetCell(xPos, y, runeValue, fg, bg)
		xPos++
	}
	return xPos, y
}

// FillWithChar Fills from x1,y1 through x2,y2 with the rune r, foreground color fg, background bg
func FillWithChar(r rune, x1, y1, x2, y2 int, fg termbox.Attribute, bg termbox.Attribute) {
	for xx := x1; xx <= x2; xx++ {
		for yx := y1; yx <= y2; yx++ {
			termbox.SetCell(xx, yx, r, fg, bg)
		}
	}
}

// DrawBorder Draw a border around the area inside x1,y1 -> x2, y2
func DrawBorder(x1, y1, x2, y2 int, fg termbox.Attribute, bg termbox.Attribute) {
	termbox.SetCell(x1, y1, '┌', fg, bg)
	FillWithChar('─', x1+1, y1, x2-1, y1, fg, bg)
	termbox.SetCell(x2, y1, '┐', fg, bg)

	FillWithChar('|', x1, y1+1, x1, y2-1, fg, bg)
	FillWithChar('|', x2, y1+1, x2, y2-1, fg, bg)

	termbox.SetCell(x1, y2, '└', fg, bg)
	FillWithChar('─', x1+1, y2, x2-1, y2, fg, bg)
	termbox.SetCell(x2, y2, '┘', fg, bg)
}

// AlignText Aligns the text txt within width characters using the specified alignment
func AlignText(txt string, width int, align TextAlignment) string {
	numSpaces := width - len(txt)
	switch align {
	case AlignCenter:
		if numSpaces/2 > 0 {
			return fmt.Sprintf("%s%s%s",
				strings.Repeat(" ", numSpaces/2),
				txt, strings.Repeat(" ", numSpaces/2),
			)
		}
		return txt
	case AlignRight:
		return fmt.Sprintf("%s%s", strings.Repeat(" ", numSpaces), txt)
	default:
		return fmt.Sprintf("%s%s", txt, strings.Repeat(" ", numSpaces))
	}
}

/* More advanced things are in their respective files */
