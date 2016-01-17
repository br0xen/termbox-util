package termboxUtil

import "github.com/nsf/termbox-go"

// ScrollFrame is a frame for holding other elements
// It manages it's own x/y, tab index
type ScrollFrame struct {
	x, y, width, height int
	scrollX, scrollY    int
	tabIdx              int
	fg, bg              termbox.Attribute
	bordered            bool
	controls            []termboxControl
}

// CreateScrollFrame creates Scrolling Frame at x, y that is w by h
func CreateScrollFrame(x, y, w, h int, fg, bg termbox.Attribute) *ScrollFrame {
	s := ScrollFrame{x: x, y: y, width: w, height: h, fg: fg, bg: bg}
	return &s
}

// GetX returns the x position of the scroll frame
func (s *ScrollFrame) GetX() int { return s.x }

// SetX sets the x position of the scroll frame
func (s *ScrollFrame) SetX(x int) {
	s.x = x
}

// GetY returns the y position of the scroll frame
func (s *ScrollFrame) GetY() int { return s.y }

// SetY sets the y position of the scroll frame
func (s *ScrollFrame) SetY(y int) {
	s.y = y
}

// GetWidth returns the current width of the scroll frame
func (s *ScrollFrame) GetWidth() int { return s.width }

// SetWidth sets the current width of the scroll frame
func (s *ScrollFrame) SetWidth(w int) {
	s.width = w
}

// GetHeight returns the current height of the scroll frame
func (s *ScrollFrame) GetHeight() int { return s.height }

// SetHeight sets the current height of the scroll frame
func (s *ScrollFrame) SetHeight(h int) {
	s.height = h
}

// IsBordered returns true or false if this scroll frame has a border
func (s *ScrollFrame) IsBordered() bool { return s.bordered }

// SetBordered sets whether we render a border around the scroll frame
func (s *ScrollFrame) SetBordered(b bool) {
	s.bordered = b
}

// GetScrollX returns the x distance scrolled
func (s *ScrollFrame) GetScrollX() int {
	return s.scrollX
}

// GetScrollY returns the y distance scrolled
func (s *ScrollFrame) GetScrollY() int {
	return s.scrollY
}

// ScrollDown scrolls the frame down
func (s *ScrollFrame) ScrollDown() {
	s.scrollY++
}

// ScrollUp scrolls the frame up
func (s *ScrollFrame) ScrollUp() {
	if s.scrollY > 0 {
		s.scrollY--
	}
}

// ScrollLeft scrolls the frame left
func (s *ScrollFrame) ScrollLeft() {
	if s.scrollX > 0 {
		s.scrollX--
	}
}

// ScrollRight scrolls the frame right
func (s *ScrollFrame) ScrollRight() {
	s.scrollX++
}

// AddControl adds a control to the frame
func (s *ScrollFrame) AddControl(t termboxControl) {
	s.controls = append(s.controls, t)
}

// DrawControl figures out the relative position of the control,
// sets it, draws it, then resets it.
func (s *ScrollFrame) DrawControl(t termboxControl) {
	if s.IsVisible(t) {
		ctlX, ctlY := t.GetX(), t.GetY()
		t.SetX((s.GetX() + ctlX))
		t.SetY((s.GetY() + ctlY))
		t.Draw()
		t.SetX(ctlX)
		t.SetY(ctlY)
	}
}

// IsVisible takes a Termbox Control and returns whether
// that control would be visible in the frame
func (s *ScrollFrame) IsVisible(t termboxControl) bool {
	// Check if any part of t should be visible
	cX, cY := t.GetX(), t.GetY()
	if cX+t.GetWidth() >= s.scrollX && cX <= s.scrollX+s.width {
		return cY+t.GetHeight() >= s.scrollY && cY <= s.scrollY+s.height
	}
	return false
}

// HandleKeyPress accepts the termbox event and returns whether it was consumed
func (s *ScrollFrame) HandleKeyPress(event termbox.Event) bool {
	return false
}

// DrawToStrings generates a slice of strings with what should
// be drawn to the screen
func (s *ScrollFrame) DrawToStrings() []string {
	return []string{}
}

// Draw outputs the Scoll Frame on the screen
func (s *ScrollFrame) Draw() {
	maxWidth := s.width
	maxHeight := s.height
	x, y := s.x, s.y
	startX := s.x
	startY := s.y
	if s.bordered {
		DrawBorder(s.x, s.y, s.x+s.width, s.y+s.height, s.fg, s.bg)
		maxWidth--
		maxHeight--
		x++
		y++
		startX++
		startY++
	}
	for i := range s.controls {
		s.DrawControl(s.controls[i])
	}
}
