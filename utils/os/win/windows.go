package win

import (
	"fmt"
	"github.com/lxn/win"
	"syscall"
	"unsafe"
)

var (
	user32DLL       = syscall.MustLoadDLL("User32.dll")
	procEnumWindows = user32DLL.MustFindProc("EnumWindows")
	findWindow      = user32DLL.MustFindProc("FindWindowW")
	getClassName    = user32DLL.MustFindProc("GetClassNameW")
	setForeground   = user32DLL.MustFindProc("SetForegroundWindow")
	getWindowRect   = user32DLL.MustFindProc("GetWindowRect")
	getWindowTextW  = user32DLL.MustFindProc("GetWindowTextW")
)

func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0)
	return &chars[0]
}

// 回调函数，用于EnumWindows中的回调函数，第一个参数是hWnd，第二个是自定义穿的参数
func AddElementFunc(hWnd win.HWND, hWndList *[]win.HWND) uintptr {
	*hWndList = append(*hWndList, hWnd)
	return 1
}

// 获取桌面下的所有窗口句柄，包括没有Windows标题的或者是窗口的。
func GetDesktopWindowHWND() []win.HWND {
	var hWndList []win.HWND
	hL := &hWndList
	_, _, err := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(syscall.NewCallback(AddElementFunc)), uintptr(unsafe.Pointer(hL)), 0)
	if err != 0 {
		fmt.Println(err)
	}
	return hWndList
}

type rect struct {
	left, top, right, bottom int32
}

func FindWindow(title, processName string) win.HWND {
	hwnd, _, _ := findWindow.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))))
	if hwnd == 0 {
		handle, _ := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
		if handle == syscall.InvalidHandle {
			return 0
		}
		defer syscall.CloseHandle(handle)
		var processEntry syscall.ProcessEntry32
		processEntry.Size = uint32(unsafe.Sizeof(processEntry))
		for syscall.Process32Next(handle, &processEntry) == nil {
			if syscall.UTF16ToString(processEntry.ExeFile[:]) == processName {
				handle, _ := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION, false, processEntry.ProcessID)
				if handle == 0 {
					return 0
				}
				defer syscall.CloseHandle(handle)
				hwnd, _, _ = findWindow.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(processName))))
				break
			}
		}
	}
	return win.HWND(hwnd)
}
func GetClassName(hwnd win.HWND) string {
	var buffer [256]uint16
	getClassName.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
	return syscall.UTF16ToString(buffer[:])
}

func GetWindowTextW(hwnd win.HWND) string {
	var buffer [256]uint16
	getWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
	return syscall.UTF16ToString(buffer[:])
}

func GetWindowRect(hwnd win.HWND) rect {
	var r rect
	getWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&r)))
	return r
}
func SetForeground(hwnd win.HWND) {
	setForeground.Call(uintptr(hwnd))
}
