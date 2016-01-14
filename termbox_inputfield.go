package termboxUtil

import "github.com/nsf/termbox-go"

// InputField is a field for inputting text
type InputField struct {
	value               string
	x, y, width, height int
	cursor              int
	fg, bg              termbox.Attribute
	bordered            bool
	wrap                bool
	multiline           bool
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
func (i *InputField) SetWrap(b bool) *InputField {
	i.wrap = b
	return i
}

// IsMultiline returns true or false if this field can have multiple lines
func (i *InputField) IsMultiline() bool { return i.multiline }

// SetMultiline sets whether the field can have multiple lines
func (i *InputField) SetMultiline(b bool) *InputField {
	i.multiline = b
	return i
}

// HandleKeyPress accepts the termbox event and returns whether it was consumed
func (i *InputField) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
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
		// Ctrl+U Clears the Input (before the cursor)
		i.value = i.value[i.cursor:]
	} else {
		// Get the rune to add to our value. Space and Tab are special cases where
		// we can't use the event's rune directly
		var ch string
		switch event.Key {
		case termbox.KeySpace:
			ch = " "
		case termbox.KeyTab:
			ch = "\t"
			/* Multiline is disabled right now
			case termbox.KeyEnter:
				if i.multiline {
					ch = "\n"
				}
			*/
		default:
			if KeyIsAlphaNumeric(event) || KeyIsSymbol(event) {
				ch = string(event.Ch)
			}
		}

		if i.cursor+len(i.value) == 0 {
			i.value = string(ch) + i.value
		} else if i.cursor == 0 {
			i.value = i.value + string(ch)
		} else {
			strPt1 := i.value[:(len(i.value) + i.cursor)]
			strPt2 := i.value[(len(i.value) + i.cursor):]
			i.value = strPt1 + string(ch) + strPt2
		}
	}
	return true
}

// Draw outputs the input field on the screen
func (i *InputField) Draw() {
	maxWidth := i.width
	maxHeight := i.height
	x, y := i.x, i.y
	startX := i.x
	startY := i.y
	if i.bordered {
		DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
		maxWidth--
		maxHeight--
		x++
		y++
		startX++
		startY++
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
	// Original:
	/*
		x, y = DrawStringAtPoint(strPt1, i.x+1, i.y+1, i.fg, i.bg)
		termbox.SetCell(x, y, cursorRune, i.bg, i.fg)
		DrawStringAtPoint(strPt2, x+1, y, i.fg, i.bg)
	*/
	if i.wrap {
		// Split the text into maxWidth chunks
		for len(strPt1) > maxWidth {
			breakAt := maxWidth
			DrawStringAtPoint(strPt1[:breakAt], x, y, i.fg, i.bg)
			x = startX
			y++
			strPt1 = strPt1[breakAt:]
		}
		x, y = DrawStringAtPoint(strPt1, x, y, i.fg, i.bg)
		if x >= maxWidth {
			y++
			x = startX
		}
		termbox.SetCell(x, y, cursorRune, i.bg, i.fg)
		x++
		if len(strPt2) > 0 {
			lenLeft := maxWidth - len(strPt1) - 1
			if lenLeft > 0 && len(strPt2) > lenLeft {
				DrawStringAtPoint(strPt2[:lenLeft], x+1, y, i.fg, i.bg)
				strPt2 = strPt2[lenLeft:]
			}
			for len(strPt2) > maxWidth {
				breakAt := maxWidth
				DrawStringAtPoint(strPt2[:breakAt], x, y, i.fg, i.bg)
				x = startX
				y++
				strPt2 = strPt2[breakAt:]
			}
			x, y = DrawStringAtPoint(strPt2, x, y, i.fg, i.bg)
		}
	} else {
		for len(strPt1)+len(strPt2)+1 > maxWidth {
			if len(strPt1) >= len(strPt2) {
				if len(strPt1) == 0 {
					break
				}
				strPt1 = strPt1[1:]
			} else {
				strPt2 = strPt2[:len(strPt2)-1]
			}
		}
		x, y = DrawStringAtPoint(strPt1, i.x+1, i.y+1, i.fg, i.bg)
		termbox.SetCell(x, y, cursorRune, i.bg, i.fg)
		DrawStringAtPoint(strPt2, x+1, y, i.fg, i.bg)
	}
}
