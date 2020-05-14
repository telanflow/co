package co

import (
	"testing"
)

func TestIsWindows(t *testing.T) {
	t.Log(IsWindows())
}

func TestIsLinux(t *testing.T) {
	t.Log(IsLinux())
}

func TestIsDarwin(t *testing.T) {
	t.Log(IsDarwin())
}

func TestGetCwd(t *testing.T) {
	t.Log(GetCwd())
}
