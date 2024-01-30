package win

import (
	"fmt"
	"github.com/gonutz/w32/v2"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
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
	_, _, err := syscall.Syscall(procEnumWindows.Addr(), 2, syscall.NewCallback(AddElementFunc), uintptr(unsafe.Pointer(hL)), 0)
	if err != 0 {
		fmt.Println(err)
	}
	return hWndList
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

func findProcess(name string) uint32 {

	processIDs, ok := w32.EnumProcesses(make([]uint32, 256))
	if !ok {
		return 0
	}

	for i := 0; i < len(processIDs); i++ {
		if processIDs[i] != 0 {
			if name == getProcName(processIDs[i]) {
				return processIDs[i]
			}
		}
	}

	return 0
}

func getProcName(pid uint32) string {

	hProcess := w32.OpenProcess(
		windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION,
		false,
		pid,
	)

	if hProcess == 0 {
		return ""
	}

	defer func() {
		w32.CloseHandle(hProcess)
	}()

	return w32.GetModuleFileName(w32.HMODULE(hProcess))
}
