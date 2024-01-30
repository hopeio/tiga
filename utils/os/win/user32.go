package win

import (
	"syscall"
)

var (
	user32DLL       = syscall.NewLazyDLL("User32.dll")
	procEnumWindows = user32DLL.NewProc("EnumWindows")
	findWindow      = user32DLL.NewProc("FindWindowW")
	getClassName    = user32DLL.NewProc("GetClassNameW")
	setForeground   = user32DLL.NewProc("SetForegroundWindow")
	getWindowRect   = user32DLL.NewProc("GetWindowRect")
	getWindowTextW  = user32DLL.NewProc("GetWindowTextW")
)
