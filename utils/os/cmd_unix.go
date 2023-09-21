//go:build unix

package osi

func ContainQuotedCMD(s string) (string, error) {
	return Cmd(s)
}
