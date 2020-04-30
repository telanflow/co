package co

import (
	"os"
	"path/filepath"
	"runtime"
)

// 当前操作系统是否Windows
func IsWin() bool {
	return "windows" == runtime.GOOS
}

// 当前操作系统是否Linux
func IsLinux() bool {
	return "linux" == runtime.GOOS
}

// 当前操作系统是否Mac OS/X
func IsMac() bool {
	return "darwin" == runtime.GOOS
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
