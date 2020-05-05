package co

import (
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
