package termbox_util

import (
	"github.com/nsf/termbox-go"
)

type ConfirmModal struct {
	title               string
	text                string
	x, y, width, height int
	show_help           bool
	cursor              int
	bg, fg              termbox.Attribute
	is_done             bool
	accepted            bool
	value               string
}

func CreateConfirmModal(title string, x, y, width, height int, fg, bg termbox.Attribute) *ConfirmModal {
	i := ConfirmModal{title: title, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	if i.title == "" && i.text == "" {
		i.title = "Confirm?"
	}
	i.show_help = true
	return &i
}

func (i *ConfirmModal) GetTitle() string { return i.title }
func (i *ConfirmModal) SetTitle(s string) *ConfirmModal {
	i.title = s
	return i
}

func (i *ConfirmModal) GetText() string { return i.text }
func (i *ConfirmModal) SetText(s string) *ConfirmModal {
	i.text = s
	return i
}

func (i *ConfirmModal) GetX() int { return i.x }
func (i *ConfirmModal) SetX(x int) *ConfirmModal {
	i.x = x
	return i
}
func (i *ConfirmModal) GetY() int { return i.y }
func (i *ConfirmModal) SetY(y int) *ConfirmModal {
	i.y = y
	return i
}

func (i *ConfirmModal) GetWidth() int { return i.width }
func (i *ConfirmModal) SetWidth(width int) *ConfirmModal {
	i.width = width
	return i
}

func (i *ConfirmModal) GetHeight() int { return i.height }
func (i *ConfirmModal) SetHeight(height int) *ConfirmModal {
	i.height = height
	return i
}

func (i *ConfirmModal) HelpIsShown() bool { return i.show_help }
func (i *ConfirmModal) ShowHelp(b bool) *ConfirmModal {
	i.show_help = b
	return i
}

func (i *ConfirmModal) GetBackground() termbox.Attribute { return i.bg }
func (i *ConfirmModal) SetBackground(bg termbox.Attribute) *ConfirmModal {
	i.bg = bg
	return i
}

func (i *ConfirmModal) GetForeground() termbox.Attribute { return i.fg }
func (i *ConfirmModal) SetForeground(fg termbox.Attribute) *ConfirmModal {
	i.fg = fg
	return i
}

func (i *ConfirmModal) IsDone() bool { return i.is_done }
func (i *ConfirmModal) SetDone(b bool) *ConfirmModal {
	i.is_done = b
	return i
}

func (i *ConfirmModal) IsAccepted() bool { return i.accepted }

func (i *ConfirmModal) Clear() *ConfirmModal {
	i.title = ""
	i.text = ""
	i.accepted = false
	i.is_done = false
	return i
}

func (i *ConfirmModal) HandleKeyPress(event termbox.Event) bool {
	if event.Ch == 'Y' || event.Ch == 'y' {
		i.accepted = true
		i.is_done = true
		return true
	} else if event.Ch == 'N' || event.Ch == 'n' {
		i.accepted = false
		i.is_done = true
		return true
	}
	return false
}
func (i *ConfirmModal) Draw() {
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
	next_y += 3
	if i.show_help {
		DrawStringAtPoint("(Y/y) Confirm. (N/n) Reject.", i.x+1, next_y, i.fg, i.bg)
	}
}
