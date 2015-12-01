package termboxUtil

import "github.com/nsf/termbox-go"

// Menu is a menu with a list of options
type Menu struct {
	title   string
	options []MenuOption
	// If height is -1, then it is adaptive to the menu
	x, y, width, height    int
	showHelp               bool
	cursor                 int
	bg, fg                 termbox.Attribute
	selectedBg, selectedFg termbox.Attribute
	disabledBg, disabledFg termbox.Attribute
	isDone                 bool
	bordered               bool
	vimMode									bool
}

// CreateMenu Creates a menu with the specified attributes
func CreateMenu(title string, options []string, x, y, width, height int, fg, bg termbox.Attribute) *Menu {
	i := Menu{
		title: title,
		x:     x, y: y, width: width, height: height,
		fg: fg, bg: bg, selectedFg: bg, selectedBg: fg,
		disabledFg: bg, disabledBg: bg,
	}
	for _, line := range options {
		i.options = append(i.options, MenuOption{text: line})
	}
	if len(i.options) > 0 {
		i.SetSelectedOption(&i.options[0])
	}
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
func (i *Menu) GetOptions() []MenuOption {
	return i.options
}

// SetOptions set the menu's options to opts
func (i *Menu) SetOptions(opts []MenuOption) *Menu {
	i.options = opts
	return i
}

// SetOptionsFromStrings sets the options of this menu from a slice of strings
func (i *Menu) SetOptionsFromStrings(opts []string) *Menu {
	var newOpts []MenuOption
	for _, v := range opts {
		newOpts = append(newOpts, *CreateOptionFromText(v))
	}
	return i.SetOptions(newOpts)
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

// GetSelectedOption returns the current selected option
func (i *Menu) GetSelectedOption() *MenuOption {
	idx := i.GetSelectedIndex()
	if idx != -1 {
		return &i.options[idx]
	}
	return nil
}

// GetOptionFromIndex Returns the
func (i *Menu) GetOptionFromIndex(idx int) *MenuOption {
	if idx >= 0 && idx < len(i.options) {
		return &i.options[idx]
	}
	return nil
}

// GetOptionFromText Returns the first option with the text v
func (i *Menu) GetOptionFromText(v string) *MenuOption {
	for idx := range i.options {
		testOption := &i.options[idx]
		if testOption.GetText() == v {
			return testOption
		}
	}
	return nil
}

// GetSelectedIndex returns the index of the selected option
// Returns -1 if nothing is selected
func (i *Menu) GetSelectedIndex() int {
	for idx := range i.options {
		if i.options[idx].IsSelected() {
			return idx
		}
	}
	return -1
}

// SetSelectedOption sets the current selected option to v (if it's valid)
func (i *Menu) SetSelectedOption(v *MenuOption) *Menu {
	for idx := range i.options {
		if &i.options[idx] == v {
			i.options[idx].Select()
		} else {
			i.options[idx].Unselect()
		}
	}
	return i
}

// SelectPrevOption Decrements the selected option (if it can)
func (i *Menu) SelectPrevOption() *Menu {
	idx := i.GetSelectedIndex()
	for idx >= 0 {
		idx--
		testOption := i.GetOptionFromIndex(idx)
		if testOption != nil && !testOption.IsDisabled() {
			i.SetSelectedOption(testOption)
			return i
		}
	}
	return i
}

// SelectNextOption Increments the selected option (if it can)
func (i *Menu) SelectNextOption() *Menu {
	idx := i.GetSelectedIndex()
	for idx < len(i.options) {
		idx++
		testOption := i.GetOptionFromIndex(idx)
		if testOption != nil && !testOption.IsDisabled() {
			i.SetSelectedOption(testOption)
			return i
		}
	}
	return i
}

// SetOptionDisabled Disables the specified option
func (i *Menu) SetOptionDisabled(idx int) {
	if len(i.options) > idx {
		i.GetOptionFromIndex(idx).Disable()
	}
}

// SetOptionEnabled Enables the specified option
func (i *Menu) SetOptionEnabled(idx int) {
	if len(i.options) > idx {
		i.GetOptionFromIndex(idx).Enable()
	}
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

// EnableVimMode Enables h,j,k,l navigation
func (i *Menu) EnableVimMode() *Menu {
	i.vimMode = true
	return i
}

// DisableVimMode Disables h,j,k,l navigation
func (i *Menu) DisableVimMode() *Menu {
	i.vimMode = false
	return i
}

// HandleKeyPress handles the termbox event and returns whether it was consumed
func (i *Menu) HandleKeyPress(event termbox.Event) bool {
	if event.Key == termbox.KeyEnter {
		i.isDone = true
		return true
	}
	currentIdx := i.GetSelectedIndex()
	switch event.Key {
	case termbox.KeyArrowUp:
		i.SelectPrevOption()
	case termbox.KeyArrowDown:
		i.SelectNextOption()
	}
	if i.vimMode {
		switch event.Ch {
		case 'j':
			i.SelectNextOption()
		case 'k':
			i.SelectPrevOption()
		}
	}
	if i.GetSelectedIndex() != currentIdx {
		return true
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
		optionWidth = i.width - 1
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
	if len(i.options) > 0 {
		for idx := range i.options {
			currOpt := &i.options[idx]
			if currOpt.IsDisabled() {
				DrawStringAtPoint(currOpt.GetText(), optionStartX, optionStartY, i.disabledFg, i.disabledBg)
			} else if i.GetSelectedOption() == currOpt {
				DrawStringAtPoint(AlignText(currOpt.GetText(), optionWidth, AlignLeft), optionStartX, optionStartY, i.selectedFg, i.selectedBg)
			} else {
				DrawStringAtPoint(currOpt.GetText(), optionStartX, optionStartY, i.fg, i.bg)
			}
			optionStartY++
		}
	}
}

/* MenuOption Struct & methods */

// MenuOption An option in the menu
type MenuOption struct {
	text     string
	selected bool
	disabled bool
	helpText string
}

// CreateOptionFromText just returns a MenuOption object
// That only has it's text value set.
func CreateOptionFromText(s string) *MenuOption {
	return &MenuOption{text: s}
}

// SetText Sets the text for this option
func (i *MenuOption) SetText(s string) *MenuOption {
	i.text = s
	return i
}

// GetText Returns the text for this option
func (i *MenuOption) GetText() string { return i.text }

// Disable Sets this option to disabled
func (i *MenuOption) Disable() *MenuOption {
	i.disabled = true
	return i
}

// Enable Sets this option to enabled
func (i *MenuOption) Enable() *MenuOption {
	i.disabled = false
	return i
}

// IsDisabled returns whether this option is enabled
func (i *MenuOption) IsDisabled() bool {
	return i.disabled
}

// IsSelected Returns whether this option is selected
func (i *MenuOption) IsSelected() bool {
	return i.selected
}

// Select Sets this option to selected
func (i *MenuOption) Select() *MenuOption {
	i.selected = true
	return i
}

// Unselect Sets this option to not selected
func (i *MenuOption) Unselect() *MenuOption {
	i.selected = false
	return i
}

// SetHelpText Sets this option's help text to s
func (i *MenuOption) SetHelpText(s string) *MenuOption {
	i.helpText = s
	return i
}

// GetHelpText Returns the help text for this option
func (i *MenuOption) GetHelpText() string { return i.helpText }
