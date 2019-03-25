package main

import (
	"os"
	"strconv"
	"syscall"

	"github.com/br0xen/termbox-util"

	"github.com/nsf/termbox-go"
)

var keepRunning bool
var initialized bool
var frame *termboxUtil.Frame

func main() {
	keepRunning = true
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	mainLoop()
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
			} else if event.Key == termbox.KeyCtrlC {
				termbox.Close()
				os.Exit(0)
			}
			keepRunning = handleEvent(event)
			if !keepRunning {
				break
			}
			layoutAndDrawScreen()
		}
		if event.Type == termbox.EventResize {
			layoutAndDrawScreen()
		}
	}
}

func layoutAndDrawScreen() {
	w, h := termbox.Size()
	if !initialized {
		fg, bg := termbox.ColorWhite, termbox.ColorBlack
		frame = termboxUtil.CreateFrame(1, 1, w-3, h-3, termbox.ColorWhite, termbox.ColorBlack)
		/*
			frame.AddControl(termboxUtil.CreateASCIIArt([]string{
				"/" + strings.Repeat("=====", 5) + "\\",
				"|" + strings.Repeat(".oOo.", 5) + "|",
				"\\" + strings.Repeat("=====", 5) + "/",
			}, 1, frame.GetBottomY()+1, fg, bg))
			frame.AddControl(termboxUtil.CreateAlertModal("AlertModal", 1, 1, w-5, 6, termbox.ColorGreen, bg))
			frame.AddControl(termboxUtil.CreateConfirmModal("ConfirmModal", 1, frame.GetBottomY()+1, w-5, 7, fg, bg))
			frame.AddControl(termboxUtil.CreateInputModal("InputModal", 1, frame.GetBottomY()+1, w-5, 7, fg, bg))
		*/
		frame.AddControl(termboxUtil.CreateDropMenu("Add Control", []string{
			"AlertModal",
			"ASCIIArt",
			"ConfirmModal",
			"DropMenu",
			"Frame",
			"InputField",
			"InputModal",
			"Label",
			"Menu",
			"ProgressBar",
			"ScrollFrame",
		},
			1, frame.GetBottomY()+1, w-5, 7, fg, bg, termbox.ColorBlack, termbox.ColorGreen))
		frame.GetLastControl().SetBordered(true)
		frame.SetActiveFlag(true)

		initialized = true
	}
	termbox.Clear(0, termbox.ColorBlack)
	drawScreen()
	termboxUtil.DrawStringAtPoint(strconv.Itoa(frame.GetBottomY()), 0, h-1, termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
}

func drawScreen() {
	frame.Draw()
}

func handleEvent(event termbox.Event) bool {
	frame.HandleEvent(event)
	for _, k := range frame.GetControls() {
		switch v := k.(type) {
		case *termboxUtil.DropMenu:
			if v.IsDone() {
			}
		case *termboxUtil.AlertModal:
			if v.IsDone() {
				v.SetText("Finished")
			}
		case *termboxUtil.ConfirmModal:
			if v.IsDone() {
				v.SetText("Finished")
			}
		}
		k.SetFgColor(termbox.ColorWhite)
	}
	frame.GetActiveControl().SetFgColor(termbox.ColorGreen)
	return true
}
