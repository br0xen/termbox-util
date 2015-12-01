package termboxUtil

import (
	"github.com/nsf/termbox-go"
)

// ConfirmModal is a modal with yes/no (or similar) buttons
type ConfirmModal struct {
	title               string
	text                string
	x, y, width, height int
	showHelp            bool
	cursor              int
	bg, fg              termbox.Attribute
	isDone              bool
	accepted            bool
	value               string
	isVisible           bool
}

// CreateConfirmModal Creates a confirmation modal with the specified attributes
func CreateConfirmModal(title string, x, y, width, height int, fg, bg termbox.Attribute) *ConfirmModal {
	i := ConfirmModal{title: title, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	if i.title == "" && i.text == "" {
		i.title = "Confirm?"
	}
	i.showHelp = true
	return &i
}

// GetTitle returns the current title of the modal
func (i *ConfirmModal) GetTitle() string { return i.title }

// SetTitle sets the current title of the modal to s
func (i *ConfirmModal) SetTitle(s string) *ConfirmModal {
	i.title = s
	return i
}

// GetText returns the current text of the modal
func (i *ConfirmModal) GetText() string { return i.text }

// SetText sets the text of the modal to s
func (i *ConfirmModal) SetText(s string) *ConfirmModal {
	i.text = s
	return i
}

// GetX returns the current x coordinate of the modal
func (i *ConfirmModal) GetX() int { return i.x }

// SetX sets the current x coordinate of the modal to x
func (i *ConfirmModal) SetX(x int) *ConfirmModal {
	i.x = x
	return i
}

// GetY returns the current y coordinate of the modal
func (i *ConfirmModal) GetY() int { return i.y }

// SetY sets the current y coordinate of the modal to y
func (i *ConfirmModal) SetY(y int) *ConfirmModal {
	i.y = y
	return i
}

// GetWidth returns the current width of the modal
func (i *ConfirmModal) GetWidth() int { return i.width }

// SetWidth sets the current modal width to width
func (i *ConfirmModal) SetWidth(width int) *ConfirmModal {
	i.width = width
	return i
}

// GetHeight returns the current height of the modal
func (i *ConfirmModal) GetHeight() int { return i.height }

// SetHeight set the height of the modal to height
func (i *ConfirmModal) SetHeight(height int) *ConfirmModal {
	i.height = height
	return i
}

// HelpIsShown returns true or false if the help is displayed
func (i *ConfirmModal) HelpIsShown() bool { return i.showHelp }

// ShowHelp sets whether or not to display the help text
func (i *ConfirmModal) ShowHelp(b bool) *ConfirmModal {
	i.showHelp = b
	return i
}

// GetBackground returns the current background color
func (i *ConfirmModal) GetBackground() termbox.Attribute { return i.bg }

// SetBackground sets the background color to bg
func (i *ConfirmModal) SetBackground(bg termbox.Attribute) *ConfirmModal {
	i.bg = bg
	return i
}

// GetForeground returns the current foreground color
func (i *ConfirmModal) GetForeground() termbox.Attribute { return i.fg }

// SetForeground sets the current foreground color to fg
func (i *ConfirmModal) SetForeground(fg termbox.Attribute) *ConfirmModal {
	i.fg = fg
	return i
}

// IsDone returns whether the user has answered the modal
func (i *ConfirmModal) IsDone() bool { return i.isDone }

// SetDone sets whether the modal has completed it's purpose
func (i *ConfirmModal) SetDone(b bool) *ConfirmModal {
	i.isDone = b
	return i
}

// Show sets the visibility flag of the modal to true
func (i *ConfirmModal) Show() *ConfirmModal {
	i.isVisible = true
	return i
}

// Hide sets the visibility flag of the modal to false
func (i *ConfirmModal) Hide() *ConfirmModal {
	i.isVisible = false
	return i
}

// IsAccepted returns whether the user accepted the modal
func (i *ConfirmModal) IsAccepted() bool { return i.accepted }

// Clear clears all of the non-positional parameters of the modal
func (i *ConfirmModal) Clear() *ConfirmModal {
	i.title = ""
	i.text = ""
	i.accepted = false
	i.isDone = false
	return i
}

// HandleKeyPress handles the termbox event and returns whether it was consumed
func (i *ConfirmModal) HandleKeyPress(event termbox.Event) bool {
	if event.Ch == 'Y' || event.Ch == 'y' {
		i.accepted = true
		i.isDone = true
		return true
	} else if event.Ch == 'N' || event.Ch == 'n' {
		i.accepted = false
		i.isDone = true
		return true
	}
	return false
}

// Draw draws the modal
func (i *ConfirmModal) Draw() {
	// First blank out the area we'll be putting the modal
	FillWithChar(' ', i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	// Now draw the border
	DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)

	nextY := i.y + 1
	// The title
	if i.title != "" {
		DrawStringAtPoint(i.title, i.x+1, nextY, i.fg, i.bg)
		nextY++
		FillWithChar('-', i.x+1, nextY, i.x+i.width-1, nextY, i.fg, i.bg)
		nextY++
	}
	if i.text != "" {
		DrawStringAtPoint(i.text, i.x+1, nextY, i.fg, i.bg)
		nextY++
	}
	nextY += 2
	if i.showHelp {
		helpString := " (Y/y) Confirm. (N/n) Reject. "
		helpX := (i.x + i.width) - len(helpString) - 1
		DrawStringAtPoint(helpString, helpX, nextY, i.fg, i.bg)
	}
}
