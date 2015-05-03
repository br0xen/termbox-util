package main

import (
	"fmt"
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

//var input *termbox_util.InputField
var input *termbox_util.InputModal
var new_key, new_val string
var mode string
var drawY int
var added_stuff []string

func layoutAndDrawScreen() {
	termbox.Clear(0, termbox.ColorBlack)
	drawScreen()
	termbox.Flush()
}

func drawScreen() {
	w, h := termbox.Size()
	termbox_util.DrawStringAtPoint(termbox_util.AlignText("Termbox Utility Test", w, termbox_util.ALIGN_CENTER), 0, 0, termbox.ColorWhite, termbox.ColorRed)
	if input == nil {
		mw, mh := w/4, h/4
		mx, my := (w/2)-(mw/2), (h/2)-(mh/2)
		input = termbox_util.CreateInputModal("", mx, my, mw, mh, termbox.ColorWhite, termbox.ColorBlack)
		input.Clear()
	}
	if mode == "bucket" {
		if input.IsDone() {
			added_stuff = append(added_stuff, fmt.Sprintf("New Bucket %s", input.GetValue()))
			input.Clear()
			mode = ""
		} else {
			input.Draw()
		}
	} else if mode == "pair" {
		if input.IsDone() {
			if new_key == "" {
				new_key = input.GetValue()
				input.Clear()
				input.SetTitle("Pair Value")
			} else {
				added_stuff = append(added_stuff, fmt.Sprintf("New Pair %s => %s", new_key, input.GetValue()))
				mode = ""
				input.Clear()
			}
		}
		if mode == "pair" && !input.IsDone() {
			input.Draw()
		}
	}
	if mode == "" {
		for i := range added_stuff {
			termbox_util.DrawStringAtPoint(added_stuff[i], 1, 3+i, termbox.ColorWhite, termbox.ColorRed)
		}
	}
}

func handleKeyEvent(event termbox.Event) bool {
	if event.Key == termbox.KeyEsc {
		return false
	} else if event.Key == termbox.KeyCtrlB {
		mode = "bucket"
		new_key = ""
		new_val = ""
		input.Clear()
		input.SetTitle("Bucket Name")
	} else if event.Key == termbox.KeyCtrlP {
		mode = "pair"
		new_key = ""
		new_val = ""
		input.Clear()
		input.SetTitle("Pair Key")
	} else {
		input.HandleKeyPress(event)
	}
	return true
}

func mainLoop() {
	added_stuff = append(added_stuff, "Ctrl+B = Add Bucket; Ctrl+P = Add Pair")
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