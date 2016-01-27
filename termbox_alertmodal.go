package termboxUtil

import (
	"github.com/nsf/termbox-go"
)

// AlertModal is a modal with yes/no (or similar) buttons
type AlertModal struct {
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

// CreateAlertModal Creates a confirmation modal with the specified attributes
func CreateAlertModal(title string, x, y, width, height int, fg, bg termbox.Attribute) *AlertModal {
	i := AlertModal{title: title, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	if i.title == "" && i.text == "" {
		i.title = "Alert!"
	}
	i.showHelp = true
	return &i
}

// GetTitle returns the current title of the modal
func (i *AlertModal) GetTitle() string { return i.title }

// SetTitle sets the current title of the modal to s
func (i *AlertModal) SetTitle(s string) {
	i.title = s
}

// GetText returns the current text of the modal
func (i *AlertModal) GetText() string { return i.text }

// SetText sets the text of the modal to s
func (i *AlertModal) SetText(s string) {
	i.text = s
}

// GetX returns the current x coordinate of the modal
func (i *AlertModal) GetX() int { return i.x }

// SetX sets the current x coordinate of the modal to x
func (i *AlertModal) SetX(x int) {
	i.x = x
}

// GetY returns the current y coordinate of the modal
func (i *AlertModal) GetY() int { return i.y }

// SetY sets the current y coordinate of the modal to y
func (i *AlertModal) SetY(y int) {
	i.y = y
}

// GetWidth returns the current width of the modal
func (i *AlertModal) GetWidth() int { return i.width }

// SetWidth sets the current modal width to width
func (i *AlertModal) SetWidth(width int) {
	i.width = width
}

// GetHeight returns the current height of the modal
func (i *AlertModal) GetHeight() int { return i.height }

// SetHeight set the height of the modal to height
func (i *AlertModal) SetHeight(height int) {
	i.height = height
}

// HelpIsShown returns true or false if the help is displayed
func (i *AlertModal) HelpIsShown() bool { return i.showHelp }

// ShowHelp sets whether or not to display the help text
func (i *AlertModal) ShowHelp(b bool) {
	i.showHelp = b
}

// GetBackground returns the current background color
func (i *AlertModal) GetBackground() termbox.Attribute { return i.bg }

// SetBackground sets the background color to bg
func (i *AlertModal) SetBackground(bg termbox.Attribute) {
	i.bg = bg
}

// GetForeground returns the current foreground color
func (i *AlertModal) GetForeground() termbox.Attribute { return i.fg }

// SetForeground sets the current foreground color to fg
func (i *AlertModal) SetForeground(fg termbox.Attribute) {
	i.fg = fg
}

// IsDone returns whether the user has answered the modal
func (i *AlertModal) IsDone() bool { return i.isDone }

// SetDone sets whether the modal has completed it's purpose
func (i *AlertModal) SetDone(b bool) {
	i.isDone = b
}

// Show sets the visibility flag of the modal to true
func (i *AlertModal) Show() {
	i.isVisible = true
}

// Hide sets the visibility flag of the modal to false
func (i *AlertModal) Hide() {
	i.isVisible = false
}

// IsAccepted returns whether the user accepted the modal
func (i *AlertModal) IsAccepted() bool { return i.accepted }

// Clear clears all of the non-positional parameters of the modal
func (i *AlertModal) Clear() {
	i.title = ""
	i.text = ""
	i.accepted = false
	i.isDone = false
}

// HandleKeyPress handles the termbox event and returns whether it was consumed
func (i *AlertModal) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		i.isDone = true
		return true
	}
	return false
}

// Draw draws the modal
func (i *AlertModal) Draw() {
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
		helpString := "Press Enter to Continue"
		helpX := (i.x + i.width) - len(helpString) - 1
		DrawStringAtPoint(helpString, helpX, nextY, i.fg, i.bg)
	}
}
