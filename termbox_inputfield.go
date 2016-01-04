package termboxUtil

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// InputField is a field for inputting text
type InputField struct {
	value               string
	x, y, width, height int
	cursor              int
	fg, bg              termbox.Attribute
	bordered            bool
	wrap                bool
}

// CreateInputField creates an input field at x, y that is w by h
func CreateInputField(x, y, w, h int, fg, bg termbox.Attribute) *InputField {
	i := InputField{x: x, y: y, width: w, height: h, fg: fg, bg: bg}
	return &i
}

// GetValue gets the current text that is in the InputField
func (i *InputField) GetValue() string { return i.value }

// SetValue sets the current text in the InputField to s
func (i *InputField) SetValue(s string) *InputField {
	i.value = s
	return i
}

// GetX returns the x position of the input field
func (i *InputField) GetX() int { return i.x }

// SetX sets the x position of the input field
func (i *InputField) SetX(x int) *InputField {
	i.x = x
	return i
}

// GetY returns the y position of the input field
func (i *InputField) GetY() int { return i.y }

// SetY sets the y position of the input field
func (i *InputField) SetY(y int) *InputField {
	i.y = y
	return i
}

// GetWidth returns the current width of the input field
func (i *InputField) GetWidth() int { return i.width }

// SetWidth sets the current width of the input field
func (i *InputField) SetWidth(w int) *InputField {
	i.width = w
	return i
}

// GetHeight returns the current height of the input field
func (i *InputField) GetHeight() int { return i.height }

// SetHeight sets the current height of the input field
func (i *InputField) SetHeight(h int) *InputField {
	i.height = h
	return i
}

// IsBordered returns true or false if this input field has a border
func (i *InputField) IsBordered() bool { return i.bordered }

// SetBordered sets whether we render a border around the input field
func (i *InputField) SetBordered(b bool) *InputField {
	i.bordered = b
	return i
}

// DoesWrap returns true or false if this input field wraps text
func (i *InputField) DoesWrap() bool { return i.wrap }

// SetWrap sets whether we wrap the text at width.
// If 'wrap' is set, we automatically increase the height when we need to.
func (i *InputField) SetWrap(b bool) *InputField {
	i.wrap = b
	return i
}

// HandleKeyPress accepts the termbox event and returns whether it was consumed
func (i *InputField) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		// Done editing
	} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
		if len(i.value) > 0 {
			i.value = i.value[:len(i.value)-1]
		}
	} else if event.Key == termbox.KeyArrowLeft {
		if i.cursor+len(i.value) > 0 {
			i.cursor--
		}
	} else if event.Key == termbox.KeyArrowRight {
		if i.cursor < 0 {
			i.cursor++
		}
	} else if event.Key == termbox.KeyCtrlU {
		// Ctrl+U Clears the Input
		i.value = ""
	} else {
		// Get the rune to add to our value. Space and Tab are special cases where
		// we can't use the event's rune directly
		var ch string
		switch event.Key {
		case termbox.KeySpace:
			ch = " "
		case termbox.KeyTab:
			ch = "\t"
		default:
			ch = string(event.Ch)
		}

		if i.cursor+len(i.value) == 0 {
			i.value = fmt.Sprintf("%s%s", ch, i.value)
		} else if i.cursor == 0 {
			i.value = fmt.Sprintf("%s%s", i.value, ch)
		} else {
			strPt1 := i.value[:(len(i.value) + i.cursor)]
			strPt2 := i.value[(len(i.value) + i.cursor):]
			i.value = fmt.Sprintf("%s%s%s", strPt1, ch, strPt2)
		}
	}
	return true
}

// Draw outputs the input field on the screen
func (i *InputField) Draw() {
	if i.bordered {
		DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	}

	var strPt1, strPt2 string
	var cursorRune rune
	if len(i.value) > 0 {
		if i.cursor+len(i.value) == 0 {
			strPt1 = ""
			strPt2 = i.value[1:]
			cursorRune = rune(i.value[0])
		} else if i.cursor == 0 {
			strPt1 = i.value
			strPt2 = ""
			cursorRune = ' '
		} else {
			strPt1 = i.value[:(len(i.value) + i.cursor)]
			strPt2 = i.value[(len(i.value)+i.cursor)+1:]
			cursorRune = rune(i.value[len(i.value)+i.cursor])
		}
	} else {
		strPt1, strPt2, cursorRune = "", "", ' '
	}
	// strPt1, strPt2 = all of the text before, after the cursor
	// cursorRune is the rune on the cursor
	// Check if the value is longer than the width
	maxWidth := i.width
	if i.bordered {
		maxWidth -= 2
	}
	cursorRune2 := cursorRune
	if len(i.value) > maxWidth {
		var chopLeft bool
		for len(strPt1)+len(strPt2)+1 > maxWidth {
			if chopLeft && len(strPt1) > 0 {
				strPt1 = strPt1[1:]
			} else if !chopLeft && len(strPt2) > 0 {
				strPt2 = strPt2[:len(strPt2)-1]
			}
			chopLeft = !chopLeft
		}
	}
	x, y := DrawStringAtPoint(strPt1, i.x+1, i.y+1, i.fg, i.bg)
	termbox.SetCell(x, y, cursorRune2, i.bg, i.fg)
	DrawStringAtPoint(strPt2, x+1, y, i.fg, i.bg)
	termbox.SetCell(x, y+1, cursorRune, i.bg, i.fg)
}
