package co

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

// 当前操作系统是否Windows
func IsWindows() bool {
	return "windows" == runtime.GOOS
}

// 当前操作系统是否Linux
func IsLinux() bool {
	return "linux" == runtime.GOOS
}

// 当前操作系统是否Mac OS/X
func IsDarwin() bool {
	return "darwin" == runtime.GOOS
}

// msys(MINGW64) env，不一定支持颜色
func IsMSys() bool {
	// "MSYSTEM=MINGW64"
	if len(os.Getenv("MSYSTEM")) > 0 {
		return true
	}

	return false
}

// check out is in stderr/stdout/stdin
func IsConsole(out io.Writer) bool {
	o, ok := out.(*os.File)
	if !ok {
		return false
	}

	fd := o.Fd()

	// fix: cannot use 'o == os.Stdout' to compare
	return fd == uintptr(syscall.Stdout) || fd == uintptr(syscall.Stdin) || fd == uintptr(syscall.Stderr)
}

// 获取当前执行文件的真实目录
// 不受 os.Chdir 函数影响
func GetCwd() string {
	var (
		dir, ex string
		err     error
	)
	ex, err = os.Executable()
	if err == nil {
		exReal, _ := filepath.EvalSymlinks(ex)
		exReal, _ = filepath.Abs(exReal)
		dir = filepath.Dir(exReal)
	}
	return dir
}
