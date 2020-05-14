package co

import (
	"os"
	"path"
	"strings"
)

// 给路径补全`/`
func PathFillSlash(path string) string {
	if path == "/" {
		return path
	}
	return path + "/"
}

// 移除路径最后的`/`
func PathRemoveSlash(path string) string {
	if len(path) > 1 {
		return strings.TrimSuffix(path, "/")
	}
	return path
}

// 将path中的反斜杠'\'替换为'/'
func PathFormSlash(old string) string {
	return path.Clean(strings.ReplaceAll(old, "\\", "/"))
}

// 路径是否存在
func PathExists(path string) bool {
	if path == "" {
		return false
	}

	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

// 判断路径是否为目录
func IsDir(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return fi.IsDir()
	}
	return false
}

// 判断路径是否为文件
func IsFile(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}
