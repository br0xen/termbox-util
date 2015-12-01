package termboxUtil

import (
	"github.com/nsf/termbox-go"
)

// InputModal A modal for text input
type InputModal struct {
	title               string
	text                string
	input               *InputField
	x, y, width, height int
	showHelp            bool
	cursor              int
	bg, fg              termbox.Attribute
	isDone              bool
	isVisible           bool
}

// CreateInputModal Create an input modal with the given attributes
func CreateInputModal(title string, x, y, width, height int, fg, bg termbox.Attribute) *InputModal {
	i := InputModal{title: title, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	i.input = CreateInputField(i.x+1, i.y+3, i.width-2, 2, i.fg, i.bg)
	i.showHelp = true
	i.input.bordered = true
	return &i
}

// GetTitle Return the title of the modal
func (i *InputModal) GetTitle() string { return i.title }

// SetTitle Sets the title of the modal to s
func (i *InputModal) SetTitle(s string) *InputModal {
	i.title = s
	return i
}

// GetText Return the text of the modal
func (i *InputModal) GetText() string { return i.text }

// SetText Set the text of the modal to s
func (i *InputModal) SetText(s string) *InputModal {
	i.text = s
	return i
}

// GetX Return the x position of the modal
func (i *InputModal) GetX() int { return i.x }

// SetX set the x position of the modal to x
func (i *InputModal) SetX(x int) *InputModal {
	i.x = x
	return i
}

// GetY Return the y position of the modal
func (i *InputModal) GetY() int { return i.y }

// SetY Set the y position of the modal to y
func (i *InputModal) SetY(y int) *InputModal {
	i.y = y
	return i
}

// GetWidth Return the width of the modal
func (i *InputModal) GetWidth() int { return i.width }

// SetWidth Set the width of the modal to width
func (i *InputModal) SetWidth(width int) *InputModal {
	i.width = width
	return i
}

// GetHeight Return the height of the modal
func (i *InputModal) GetHeight() int { return i.height }

// SetHeight Set the height of the modal to height
func (i *InputModal) SetHeight(height int) *InputModal {
	i.height = height
	return i
}

// HelpIsShown Returns whether the modal is showing it's help text or not
func (i *InputModal) HelpIsShown() bool { return i.showHelp }

// ShowHelp Set the "Show Help" flag
func (i *InputModal) ShowHelp(b bool) *InputModal {
	i.showHelp = b
	return i
}

// GetBackground Return the current background color of the modal
func (i *InputModal) GetBackground() termbox.Attribute { return i.bg }

// SetBackground Set the current background color to bg
func (i *InputModal) SetBackground(bg termbox.Attribute) *InputModal {
	i.bg = bg
	return i
}

// GetForeground Return the current foreground color
func (i *InputModal) GetForeground() termbox.Attribute { return i.fg }

// SetForeground Set the foreground color to fg
func (i *InputModal) SetForeground(fg termbox.Attribute) *InputModal {
	i.fg = fg
	return i
}

// Show Sets the visibility flag to true
func (i *InputModal) Show() *InputModal {
	i.isVisible = true
	return i
}

// Hide Sets the visibility flag to false
func (i *InputModal) Hide() *InputModal {
	i.isVisible = false
	return i
}

// SetDone Sets the flag that tells whether this modal has completed it's purpose
func (i *InputModal) SetDone(b bool) *InputModal {
	i.isDone = b
	return i
}

// IsDone Returns the "isDone" flag
func (i *InputModal) IsDone() bool {
	return i.isDone
}

// GetValue Return the current value of the input
func (i *InputModal) GetValue() string { return i.input.GetValue() }

// SetValue Sets the value of the input to s
func (i *InputModal) SetValue(s string) *InputModal {
	i.input.SetValue(s)
	return i
}

// Clear Resets all non-positional parameters of the modal
func (i *InputModal) Clear() *InputModal {
	i.title = ""
	i.text = ""
	i.input.SetValue("")
	i.isDone = false
	i.isVisible = false
	return i
}

// HandleKeyPress Handle the termbox event, return true if it was consumed
func (i *InputModal) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		// Done editing
		i.isDone = true
		return true
	}
	return i.input.HandleKeyPress(event)
}

// Draw Draw the modal
func (i *InputModal) Draw() {
	if i.isVisible {
		// First blank out the area we'll be putting the modal
		FillWithChar(' ', i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
		nextY := i.y + 1
		// The title
		if i.title != "" {
			if len(i.title) > i.width {
				diff := i.width - len(i.title)
				DrawStringAtPoint(i.title[:len(i.title)+diff-1], i.x+1, nextY, i.fg, i.bg)
			} else {
				DrawStringAtPoint(i.title, i.x+1, nextY, i.fg, i.bg)
			}
			nextY++
			FillWithChar('-', i.x+1, nextY, i.x+i.width-1, nextY, i.fg, i.bg)
			nextY++
		}
		if i.text != "" {
			DrawStringAtPoint(i.text, i.x+1, nextY, i.fg, i.bg)
			nextY++
		}
		i.input.SetY(nextY)
		i.input.Draw()
		nextY += 3
		if i.showHelp {
			helpString := " (ENTER) to Accept. (ESC) to Cancel. "
			helpX := (i.x + i.width - len(helpString)) - 1
			DrawStringAtPoint(helpString, helpX, nextY, i.fg, i.bg)
		}
		// Now draw the border
		DrawBorder(i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	}
}
