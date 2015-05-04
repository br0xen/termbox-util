package termbox_util

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
)

type TextAlignment int

const (
	ALIGN_LEFT = iota
	ALIGN_CENTER
	ALIGN_RIGHT
)

/* Basic Output Helpers */
func DrawStringAtPoint(str string, x int, y int, fg termbox.Attribute, bg termbox.Attribute) (int, int) {
	x_pos := x
	for _, runeValue := range str {
		termbox.SetCell(x_pos, y, runeValue, fg, bg)
		x_pos++
	}
	return x_pos, y
}

func FillWithChar(r rune, x1, y1, x2, y2 int, fg termbox.Attribute, bg termbox.Attribute) {
	for xx := x1; xx <= x2; xx++ {
		for yx := y1; yx <= y2; yx++ {
			termbox.SetCell(xx, yx, r, fg, bg)
		}
	}
}

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

func AlignText(txt string, width int, align TextAlignment) string {
	num_spaces := width - len(txt)
	switch align {
	case ALIGN_CENTER:
		return fmt.Sprintf("%s%s%s",
			strings.Repeat(" ", num_spaces/2),
			txt, strings.Repeat(" ", num_spaces/2),
		)
	case ALIGN_RIGHT:
		return fmt.Sprintf("%s%s", strings.Repeat(" ", num_spaces), txt)
	default:
		return fmt.Sprintf("%s%s", txt, strings.Repeat(" ", num_spaces))
	}
}

/* More advanced things are in their respective files*/
