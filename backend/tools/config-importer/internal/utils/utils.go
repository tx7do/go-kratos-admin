package utils

import (
	"os"
	"path/filepath"
)

// IsExistString 判断目标字符串是否是在切片中
func IsExistString(ignores []string, str string) bool {
	if len(ignores) == 0 {
		return false
	}

	for _, ignore := range ignores {
		if ignore == str {
			return true
		}
	}

	return false
}

// GetFileList 获取文件夹下面的所有文件的列表
func GetFileList(root string) []string {
	var files []string

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	}); err != nil {
		return nil
	}

	return files
}

// GetFolderNameList 获取当前文件夹下面的所有文件夹名的列表
func GetFolderNameList(root string) []string {
	var names []string
	fs, _ := os.ReadDir(root)
	for _, file := range fs {
		if file.IsDir() {
			names = append(names, file.Name())
		}
	}
	return names
}

// ReadFile 读取文件
func ReadFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return content
}
