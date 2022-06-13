package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	keystate = user32.NewProc("GetAsyncKeyState")
)

func run(file string) {
	s := fmt.Sprintf("running %s ", file)
	fmt.Println(s)
	exec.Command(file).Start()
}

func keyLogger() {
	alacritty := "C:/Program Files/Alacritty/alacritty.exe"
	enterDown := false
	enterKey := 0x0D
	winKey := 0x5B
	for {

		enterState, _, _ := keystate.Call(uintptr(enterKey))
		winState, _, _ := keystate.Call(uintptr(winKey))

		if winState > 1 {
			if enterState > 1 {
				if !enterDown {
					run(alacritty)
					enterDown = true
				}
			} else {
				enterDown = false
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	go keyLogger()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
