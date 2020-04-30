package co

import (
	"testing"
)

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

func TestGetCwd(t *testing.T) {
	t.Log(GetCwd())
}