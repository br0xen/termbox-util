package termbox_util

import (
	"github.com/nsf/termbox-go"
)

type InputModal struct {
	title               string
	text                string
	input               *InputField
	x, y, width, height int
	show_help           bool
	cursor              int
	bg, fg              termbox.Attribute
	is_done             bool
	value               string
}

func CreateInputModal(title string, x, y, width, height int, fg, bg termbox.Attribute) *InputModal {
	i := InputModal{title: title, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	i.input = CreateInputField(i.x+1, i.y+3, i.width-2, 2, i.fg, i.bg)
	i.input.bordered = true
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
func (i *InputModal) SetDone(b bool) *InputModal {
	i.is_done = b
	return i
}
func (i *InputModal) IsDone() bool {
	return i.is_done
}
func (i *InputModal) SetValue(s string) *InputModal {
	i.input.SetValue(s)
	return i
}
func (i *InputModal) GetValue() string {
	return i.input.GetValue()
}
func (i *InputModal) Clear() *InputModal {
	i.title = ""
	i.text = ""
	i.input.SetValue("")
	i.is_done = false
	return i
}

func (i *InputModal) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		// Done editing
		i.value = i.input.GetValue()
		i.is_done = true
		return true
	} else {
		return i.input.HandleKeyPress(event)
	}
}
func (i *InputModal) Draw() {
	// First blank out the area we'll be putting the modal
	FillWithChar(' ', i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	// Now draw the border
	DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)

	next_y := i.y + 1
	// The title
	if i.title != "" {
		DrawStringAtPoint(i.title, i.x+1, next_y, i.fg, i.bg)
		next_y += 1
		FillWithChar('-', i.x+1, next_y, i.x+i.width-1, next_y, i.fg, i.bg)
		next_y += 1
	}
	if i.text != "" {
		DrawStringAtPoint(i.text, i.x+1, next_y, i.fg, i.bg)
		next_y += 1
	}
	i.input.SetY(next_y)
	i.input.Draw()
	next_y += 3
}
