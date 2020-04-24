package os

import "testing"

func TestIsWin(t *testing.T) {
	is := IsWin()
	t.Log(is)
}

func TestIsLinux(t *testing.T) {
	is := IsLinux()
	t.Log(is)
}

func TestIsMac(t *testing.T) {
	is := IsMac()
	t.Log(is)
}

func TestGetCurrentDirectory(t *testing.T) {
	path := GetCurrentDirectory()
	t.Log(path)
}

func TestPwd(t *testing.T) {
	path := Pwd()
	t.Log(path)
}