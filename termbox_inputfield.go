package termbox_util

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type InputField struct {
	value               string
	x, y, width, height int
	cursor              int
	fg, bg              termbox.Attribute
	bordered            bool
}

func CreateInputField(x, y, w, h int, fg, bg termbox.Attribute) *InputField {
	i := InputField{x: x, y: y, width: w, height: h, fg: fg, bg: bg}
	return &i
}

func (i *InputField) GetValue() string { return i.value }
func (i *InputField) SetValue(s string) *InputField {
	i.value = s
	return i
}

func (i *InputField) GetX() int { return i.x }
func (i *InputField) SetX(x int) *InputField {
	i.x = x
	return i
}

func (i *InputField) GetY() int { return i.y }
func (i *InputField) SetY(y int) *InputField {
	i.y = y
	return i
}

func (i *InputField) GetWidth() int { return i.width }
func (i *InputField) SetWidth(w int) *InputField {
	i.width = w
	return i
}

func (i *InputField) GetHeight() int { return i.height }
func (i *InputField) SetHeight(h int) *InputField {
	i.height = h
	return i
}

func (i *InputField) IsBordered() bool { return i.bordered }
func (i *InputField) SetBordered(b bool) *InputField {
	i.bordered = b
	return i
}

func (i *InputField) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		// Done editing
	} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
		if len(i.value) > 0 {
			i.value = i.value[:len(i.value)-1]
		}
	} else if event.Key == termbox.KeyArrowLeft {
		if i.cursor+len(i.value) > 0 {
			i.cursor -= 1
		}
	} else if event.Key == termbox.KeyArrowRight {
		if i.cursor < 0 {
			i.cursor += 1
		}
	} else if event.Key == termbox.KeyCtrlU {
		// Ctrl+U Clears the Input
		i.value = ""
	} else {
		if i.cursor+len(i.value) == 0 {
			i.value = fmt.Sprintf("%s%s", string(event.Ch), i.value)
		} else if i.cursor == 0 {
			i.value = fmt.Sprintf("%s%s", i.value, string(event.Ch))
		} else {
			str_pt_1 := i.value[:(len(i.value) + i.cursor)]
			str_pt_2 := i.value[(len(i.value) + i.cursor):]
			i.value = fmt.Sprintf("%s%s%s", str_pt_1, string(event.Ch), str_pt_2)
		}
	}
	return true
}

func (i *InputField) Draw() {
	if i.bordered {
		DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	}

	var str_pt_1, str_pt_2 string
	var cursor_rune rune
	if len(i.value) > 0 {
		if i.cursor+len(i.value) == 0 {
			str_pt_1 = ""
			str_pt_2 = i.value[1:]
			cursor_rune = rune(i.value[0])
		} else if i.cursor == 0 {
			str_pt_1 = i.value
			str_pt_2 = ""
			cursor_rune = ' '
		} else {
			str_pt_1 = i.value[:(len(i.value) + i.cursor)]
			str_pt_2 = i.value[(len(i.value)+i.cursor)+1:]
			cursor_rune = rune(i.value[len(i.value)+i.cursor])
		}
	} else {
		str_pt_1, str_pt_2, cursor_rune = "", "", ' '
	}
	x, y := DrawStringAtPoint(str_pt_1, i.x+1, i.y+1, i.fg, i.bg)
	termbox.SetCell(x, y, cursor_rune, i.bg, i.fg)
	DrawStringAtPoint(str_pt_2, x+1, y, i.fg, i.bg)
}
