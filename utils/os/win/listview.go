package win

import (
	"github.com/gonutz/w32/v2"
	"github.com/sirupsen/logrus"
	"syscall"
	"unsafe"
)

func listViews(hwnd w32.HWND) []w32.HWND {
	var listViewHwnds []w32.HWND
	fnOfEnumListView := func(childHwnd w32.HWND) bool {
		className, _ := w32.GetClassName(childHwnd)

		if className == "SysListView" {
			listViewHwnds = append(listViewHwnds, childHwnd)
		}
		return true
	}

	w32.EnumChildWindows(hwnd, fnOfEnumListView)

	return listViewHwnds
}

func getLVItemRowCount(hwnd w32.HWND) int {
	rowCount := w32.SendMessage(hwnd, w32.LVM_GETITEMCOUNT, 0, 0)
	return int(rowCount)
}

func getLVItem(hwnd w32.HWND, row, col int) string {

	rowCount := w32.SendMessage(hwnd, w32.LVM_GETITEMCOUNT, 0, 0)
	if rowCount == 0 {
		return ""
	}

	if row-1 > int(rowCount) {
		return ""
	}

	_, pid := w32.GetWindowThreadProcessId(hwnd)

	hProcess := w32.OpenProcess(
		w32.PROCESS_VM_READ|w32.PROCESS_VM_WRITE|w32.PROCESS_VM_OPERATION|w32.PROCESS_QUERY_INFORMATION,
		false,
		uint32(pid),
	)

	if hProcess == 0 {
		logrus.Errorln("开启远程hProcess失败")
		return ""
	}

	defer func() {
		w32.CloseHandle(hProcess)
	}()

	lpLvItem := VirtualAllocEx(hProcess, 0, unsafe.Sizeof(w32.LVITEM{}), MEM_COMMIT, PAGE_READWRITE)
	if lpLvItem == 0 {
		logrus.Errorln("申请远程内存空间失败")
		return ""
	}

	defer func() {
		VirtualFreeEx(hProcess, lpLvItem, 0, MEM_RELEASE)
	}()

	lpStr := VirtualAllocEx(hProcess, 0, 256, MEM_COMMIT, PAGE_READWRITE)
	if lpStr == 0 {
		logrus.Errorln("申请远程内存空间失败")
		return ""
	}

	defer func() {
		VirtualFreeEx(hProcess, lpStr, 0, MEM_RELEASE)
	}()

	item := &w32.LVITEM{
		Mask:       w32.LVIF_TEXT,
		IItem:      int32(row),
		ISubItem:   int32(col),
		PszText:    (*uint16)(unsafe.Pointer(lpStr)),
		CchTextMax: 256,
	}

	_, ok := WriteProcessMemory(hProcess, lpLvItem, uintptr(unsafe.Pointer(item)), unsafe.Sizeof(w32.LVITEM{}))
	if !ok {
		return ""
	}

	ret := w32.SendMessage(hwnd, w32.LVM_GETITEMTEXT, uintptr(row), lpLvItem)
	if int(ret) > 0 {
		redBuf, _, _ := ReadProcessMemory(hProcess, lpStr, ret*2)
		s := syscall.UTF16ToString(redBuf)
		return s
	}

	return ""
}
