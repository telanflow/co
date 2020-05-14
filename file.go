package co

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	// sniff Length, use for detect file mime type
	MimeSniffLen = 512
)

// 文件是否存在
func FileExists(path string) bool {
	return PathExists(path)
}

// 给定path创建文件，如果目录不存在就递归创建
func FileCreateMust(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !FileExists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			return nil, err
		}
	}

	return os.Create(path)
}

// 获取文件mime类型
// Usage:
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return
// 	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) {
	var buf [MimeSniffLen]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}

	return http.DetectContentType(buf[:n])
}

// 获取文件mime类型. eg "image/png"
func FileMimeType(path string) (mime string) {
	if path == "" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}

	return ReaderMimeType(file)
}

// 获取目录所有文件夹
func GetDirs(path string) (re []string) {
	if PathExists(path) {
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// 获取目录所有文件
func GetFiles(path string) (re []string) {
	if PathExists(path) {
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// 获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤
// 支持 Linux/macOS 软链接
func GetFilesSuffix(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 32)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	var walkFunc filepath.WalkFunc
	walkFunc = func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil {
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if fi.Mode()&os.ModeSymlink != 0 { // 读取 symbol link
			err = filepath.Walk(filename+string(os.PathSeparator), walkFunc)
			return err
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, path.Clean(filename))
		}
		return nil
	}

	err = filepath.Walk(dirPth, walkFunc)
	return files, err
}

// 按行读取文件
func ReadLine(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	list := make([]string, 0)
	bfRd := bufio.NewReader(file)
	for {
		line, err := bfRd.ReadBytes('\n')
		// 遇到任何错误立即返回，并忽略 EOF 错误信息
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		list = append(list, string(line))
	}
	return list, nil
}
