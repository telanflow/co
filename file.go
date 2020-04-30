package co

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// 文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		return true
	}
	return false
}

// 创建目录 生成多级目录
func BuildDir(dir string) error {
	return os.MkdirAll(path.Dir(dir), os.ModePerm)
}

// 删除文件或文件夹
func DeleteFile(dir string) error {
	return os.RemoveAll(dir)
}

// 获取目录所有文件夹
func GetPathDirs(dir string) (re []string) {
	if PathExists(dir) {
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
	if PathExists(dir) {
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
