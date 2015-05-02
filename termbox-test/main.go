package main

import (
	"github.com/nsf/termbox-go"
	"gogs.bullercodeworks.com/brian/termbox-util"
	"os"
	"syscall"
)

var keep_running bool

func main() {
	keep_running = true
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	mainLoop()
}

var input *termbox_util.InputField

//var input *termbox_util.InputModal

func layoutAndDrawScreen() {
	termbox.Clear(0, termbox.ColorBlack)
	drawScreen()
	termbox.Flush()
}

func drawScreen() {
	w, h := termbox.Size()
	termbox_util.DrawStringAtPoint(termbox_util.AlignText("Termbox Utility Test", w, termbox_util.ALIGN_CENTER), 0, 0, termbox.ColorWhite, termbox.ColorRed)
	if input == nil {
		//		mw, mh := w/4, h/4
		//		mx, my := w-(mw/2), h-(mh/2)
		mw, mh := w/4, 2
		mx, my := (w/2)-(mw/2), (h/2)-(mh/2)
		//		input = termbox_util.CreateInputModal("Test Input", mx, my, mw, mh, termbox.ColorWhite, termbox.ColorBlack)
		input = termbox_util.CreateInputField(mx, my, mw, mh, termbox.ColorWhite, termbox.ColorBlack)
		input.SetBordered(true)
	}
	input.Draw()
}

func handleKeyEvent(event termbox.Event) bool {
	if event.Key == termbox.KeyEsc {
		return false
	} else {
		input.HandleKeyPress(event)
	}
	return true
}

func mainLoop() {
	layoutAndDrawScreen()
	for {
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyCtrlZ {
				process, _ := os.FindProcess(os.Getpid())
				termbox.Close()
				process.Signal(syscall.SIGSTOP)
				termbox.Init()
			}
			keep_running = handleKeyEvent(event)
			if !keep_running {
				break
			}
			layoutAndDrawScreen()
		}
		if event.Type == termbox.EventResize {
			layoutAndDrawScreen()
		}
	}
}