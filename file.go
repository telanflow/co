package co

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

// 给定path创建文件，如果目录不存在就递归创建
func CreatFileNested(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !FileExists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			return nil, err
		}
	}

	return os.Create(path)
}

// 获取目录所有文件夹
func GetPathDirs(dir string) (re []string) {
	if FileExists(dir) {
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// 获取目录所有文件
func GetPathFiles(dir string) (re []string) {
	if FileExists(dir) {
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

// 按行读取文件
func ReadLine(filePath string, callback func([]byte) bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	bfRd := bufio.NewReader(file)
	for {
		line, err := bfRd.ReadBytes('\n')
		if err != nil {
			// 遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}

		if !callback(line) {
			// 返回 false 停止读取
			break
		}
	}
	return nil
}
