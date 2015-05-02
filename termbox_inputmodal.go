package termbox_util

import (
	"github.com/nsf/termbox-go"
)

type InputModal struct {
	title               string
	text                string
	value               string
	x, y, width, height int
	show_help           bool
	cursor              int
	bg, fg              termbox.Attribute
}

func CreateInputModal(text string, x, y, width, height int, fg, bg termbox.Attribute) *InputModal {
	i := InputModal{text: text, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	return &i
}

func (i *InputModal) GetTitle() string { return i.title }
func (i *InputModal) SetTitle(s string) *InputModal {
	i.title = s
	return i
}

func (i *InputModal) GetText() string { return i.text }
func (i *InputModal) SetText(s string) *InputModal {
	i.text = s
	return i
}

func (i *InputModal) GetValue() string { return i.value }
func (i *InputModal) SetValue(s string) *InputModal {
	i.value = s
	return i
}

func (i *InputModal) GetX() int { return i.x }
func (i *InputModal) SetX(x int) *InputModal {
	i.x = x
	return i
}
func (i *InputModal) GetY() int { return i.y }
func (i *InputModal) SetY(y int) *InputModal {
	i.y = y
	return i
}

func (i *InputModal) GetWidth() int { return i.width }
func (i *InputModal) SetWidth(width int) *InputModal {
	i.width = width
	return i
}

func (i *InputModal) GetHeight() int { return i.height }
func (i *InputModal) SetHeight(height int) *InputModal {
	i.height = height
	return i
}

func (i *InputModal) HelpIsShown() bool { return i.show_help }
func (i *InputModal) ShowHelp(b bool) *InputModal {
	i.show_help = b
	return i
}

func (i *InputModal) GetCursorPos() int { return i.cursor }
func (i *InputModal) SetCursorPos(c int) *InputModal {
	i.cursor = c
	return i
}
func (i *InputModal) MoveCursorLeft() *InputModal {
	if len(i.value)+(i.GetCursorPos()) > 0 {
		i.cursor = i.GetCursorPos() - 1
	}
	return i
}
func (i *InputModal) MoveCursorRight() *InputModal {
	if i.GetCursorPos() < 0 {
		i.cursor = i.GetCursorPos() + 1
	}
	return i
}

func (i *InputModal) GetBackground() termbox.Attribute { return i.bg }
func (i *InputModal) SetBackground(bg termbox.Attribute) *InputModal {
	i.bg = bg
	return i
}

func (i *InputModal) GetForeground() termbox.Attribute { return i.fg }
func (i *InputModal) SetForeground(fg termbox.Attribute) *InputModal {
	i.fg = fg
	return i
}

func (i *InputModal) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		// Done editing
	} else if event.Key == termbox.KeyBackspace || event.Key == termbox.KeyBackspace2 {
		i.value = i.value[:len(i.value)-1]
		i.cursor -= 1
	} else {
		i.value += string(event.Ch)
		i.cursor += 1
	}
	return true
}
func (i *InputModal) Draw() {
	// First blank out the area we'll be putting the modal
	FillWithChar(' ', i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	// Now draw the border
	DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)

	DrawBorder(i.x+2, i.y+2, i.x+i.width-2, i.y+4, i.fg, i.bg)
	// TODO: Output Cursor at appropriate spot
	var output_string_1, output_string_2 string
	var cursor_rune rune
	if len(i.value) > 0 {
		output_string_1 = i.value[:(len(i.value) - 1 + i.cursor)]
		output_string_2 = i.value[(len(i.value) - 1 + i.cursor):]
		cursor_rune = ' '
	} else {
		output_string_1 = ""
		output_string_2 = ""
		cursor_rune = ' '
	}

	DrawStringAtPoint(output_string_1, i.x+3, i.y+3, i.fg, i.bg)
	termbox.SetCell(i.x+len(output_string_1), i.y+3, cursor_rune, i.bg, i.fg)
	DrawStringAtPoint(output_string_2, i.x+3+len(output_string_1)+1, i.y+3, i.fg, i.bg)
}
