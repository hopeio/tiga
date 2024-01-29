package win

import (
	"fmt"
	"testing"
)

func TestWindows(t *testing.T) {
	var windows = GetDesktopWindowHWND()
	for _, w := range windows {
		fmt.Println(GetWindowTextW(w))
	}
}
