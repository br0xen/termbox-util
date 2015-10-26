package termboxUtil

import (
	"github.com/nsf/termbox-go"
)

// Menu is a menu with a list of options
type Menu struct {
	title   string
	options []string
	// If height is -1, then it is adaptive to the menu
	x, y, width, height    int
	optionsDisabled        []bool
	optionsHelp            []string
	showHelp               bool
	cursor                 int
	bg, fg                 termbox.Attribute
	selectedBg, selectedFg termbox.Attribute
	disabledBg, disabledFg termbox.Attribute
	isDone                 bool
	selectedOption         int
	bordered               bool
	hasFocus               bool
}

// CreateMenu Creates a menu with the specified attributes
func CreateMenu(title string, options []string, x, y, width, height int, fg, bg termbox.Attribute) *Menu {
	i := Menu{title: title, options: options, x: x, y: y, width: width, height: height, fg: fg, bg: bg}
	for len(i.optionsDisabled) < len(i.options) {
		i.optionsDisabled = append(i.optionsDisabled, false)
	}
	i.selectedFg = i.bg
	i.selectedBg = i.fg
	i.disabledFg = i.bg
	i.disabledBg = i.bg
	return &i
}

// GetTitle returns the current title of the menu
func (i *Menu) GetTitle() string { return i.title }

// SetTitle sets the current title of the menu to s
func (i *Menu) SetTitle(s string) *Menu {
	i.title = s
	return i
}

// GetOptions returns the current options of the menu
func (i *Menu) GetOptions() []string {
	return i.options
}

// SetOptions set the menu's options to opts
func (i *Menu) SetOptions(opts []string) *Menu {
	i.options = opts
	return i
}

// SetOptionDisabled sets an option in the menu to disabled
func (i *Menu) SetOptionDisabled(idx int) *Menu {
	if idx > 0 && idx < len(i.options) {
		i.optionsDisabled[idx] = true
	}
	return i
}

// SetOptionEnabled sets an option to enabled
func (i *Menu) SetOptionEnabled(idx int) *Menu {
	if idx >= 0 && idx < len(i.options) {
		i.optionsDisabled[idx] = false
	}
	return i
}

// GetX returns the current x coordinate of the menu
func (i *Menu) GetX() int { return i.x }

// SetX sets the current x coordinate of the menu to x
func (i *Menu) SetX(x int) *Menu {
	i.x = x
	return i
}

// GetY returns the current y coordinate of the menu
func (i *Menu) GetY() int { return i.y }

// SetY sets the current y coordinate of the menu to y
func (i *Menu) SetY(y int) *Menu {
	i.y = y
	return i
}

// GetWidth returns the current width of the menu
func (i *Menu) GetWidth() int { return i.width }

// SetWidth sets the current menu width to width
func (i *Menu) SetWidth(width int) *Menu {
	i.width = width
	return i
}

// GetHeight returns the current height of the menu
func (i *Menu) GetHeight() int { return i.height }

// SetHeight set the height of the menu to height
func (i *Menu) SetHeight(height int) *Menu {
	i.height = height
	return i
}

// HelpIsShown returns true or false if the help is displayed
func (i *Menu) HelpIsShown() bool { return i.showHelp }

// ShowHelp sets whether or not to display the help text
func (i *Menu) ShowHelp(b bool) *Menu {
	i.showHelp = b
	return i
}

// GetBackground returns the current background color
func (i *Menu) GetBackground() termbox.Attribute { return i.bg }

// SetBackground sets the background color to bg
func (i *Menu) SetBackground(bg termbox.Attribute) *Menu {
	i.bg = bg
	return i
}

// GetForeground returns the current foreground color
func (i *Menu) GetForeground() termbox.Attribute { return i.fg }

// SetForeground sets the current foreground color to fg
func (i *Menu) SetForeground(fg termbox.Attribute) *Menu {
	i.fg = fg
	return i
}

// IsDone returns whether the user has answered the modal
func (i *Menu) IsDone() bool { return i.isDone }

// SetDone sets whether the modal has completed it's purpose
func (i *Menu) SetDone(b bool) *Menu {
	i.isDone = b
	return i
}

// IsBordered returns true or false if this menu has a border
func (i *Menu) IsBordered() bool { return i.bordered }

// SetBordered sets whether we render a border around the menu
func (i *Menu) SetBordered(b bool) *Menu {
	i.bordered = b
	return i
}

// HandleKeyPress handles the termbox event and returns whether it was consumed
func (i *Menu) HandleKeyPress(event termbox.Event) bool {
	if i.hasFocus {
		if event.Key == termbox.KeyEnter {
			i.isDone = true
			return true
		}
		switch event.Key {
		case termbox.KeyArrowUp:
			if i.selectedOption > 0 {
				i.selectedOption--
				return true
			}
		case termbox.KeyArrowDown:
			if i.selectedOption < len(i.options) {
				i.selectedOption++
				return true
			}
		}
	}
	return false
}

// Draw draws the modal
func (i *Menu) Draw() {
	// First blank out the area we'll be putting the menu
	FillWithChar(' ', i.x, i.y, i.x+i.width, i.y+i.height, i.fg, i.bg)
	// Now draw the border
	optionStartX := i.x
	optionStartY := i.y
	optionWidth := i.width
	optionHeight := i.height
	if optionHeight == -1 {
		optionHeight = len(i.options)
	}
	if i.bordered {
		if i.height == -1 {
			DrawBorder(i.x, i.y, i.x+i.width, i.y+optionHeight+1, i.fg, i.bg)
		} else {
			DrawBorder(i.x, i.y, i.x+i.width, i.y+optionHeight, i.fg, i.bg)
		}
		optionStartX = i.x + 1
		optionStartY = i.y + 1
		optionWidth = i.width - 2
	}

	// The title
	if i.title != "" {
		DrawStringAtPoint(AlignText(i.title, optionWidth, AlignCenter), optionStartX, optionStartY, i.fg, i.bg)
		optionStartY++
		if i.bordered {
			FillWithChar('-', optionStartX, optionStartY, optionWidth, optionStartY, i.fg, i.bg)
			optionStartY++
		}
	}

	// Print the options
	for idx, opt := range i.options {
		if i.optionsDisabled[idx] {
			DrawStringAtPoint(opt, optionStartX, optionStartY, i.disabledFg, i.disabledBg)
		} else if i.selectedOption == idx {
			DrawStringAtPoint(opt, optionStartX, optionStartY, i.selectedFg, i.selectedBg)
		} else {
			DrawStringAtPoint(opt, optionStartX, optionStartY, i.fg, i.bg)
		}
	}
}
