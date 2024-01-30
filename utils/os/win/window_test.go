package win

import (
	"fmt"
	"github.com/gonutz/w32/v2"
	"testing"
)

func TestWindows(t *testing.T) {
	/*	w32.EnumWindows(func(w w32.HWND) bool {
			name, _ := w32.GetClassName(w)
			fmt.Println(w32.GetWindowText(w), name)
			return true
		})
	*/
	processIDs, ok := w32.EnumProcesses(make([]uint32, 1024))
	if !ok {
		return
	}

	for i := 0; i < len(processIDs); i++ {
		if processIDs[i] != 0 {
			fmt.Println(getProcName(processIDs[i]))
		}
	}

	/*	mainFormTitle := "任务管理器"
		mainFormClass := "TaskManagerWindow"
		hwnd := w32.FindWindow(mainFormClass, mainFormTitle)
		w32.ShowWindow(hwnd, w32.SW_NORMAL)
		w32.SetForegroundWindow(hwnd)*/
	/*var windows = GetDesktopWindowHWND()
	for _, w := range windows {

	}*/
}
