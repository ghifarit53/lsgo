package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func getFileExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	extension = strings.TrimPrefix(extension, ".")
	extension = strings.ToLower(extension)
	return extension
}

func isFileReadable(fileName string) bool {
	// Check whether a file/directory is readable
	_, err := os.Open(fileName)
	return err == nil
}

func isDirectory(fileName string) bool {
	// Check whether filename is a directory
	stat, err := os.Lstat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return stat.Mode().IsDir()
}

func isDirectoryEmpty(fileName string) bool {
	// Check whether directory is empty
	// https://stackoverflow.com/questions/30697324/how-to-check-if-directory-on-path-is-empty
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.Readdirnames(1)
	return err == io.EOF
}

func isFileEmpty(fileName string) bool {
	// Check whether file is empty
	// https://www.kelche.co/blog/go/golang-file-handling
	stat, err := os.Lstat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return stat.Size() == 0
}

func isFileExecutable(fileName string) bool {
	// Check whether file is executable
	// https://stackoverflow.com/questions/45429210/how-do-i-check-a-files-permissions-in-linux-using-go
	stat, err := os.Lstat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return stat.Mode()&0100 != 0
}

func isSymlink(fileName string) bool {
	// Check whether file is a symlink
	// https://stackoverflow.com/questions/64857275/detecting-symbolic-link
	stat, err := os.Lstat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return stat.Mode()&os.ModeSymlink == os.ModeSymlink
}

func isTextFile(fileName string) bool {
	// Check whether file is a textfile
	// https://stackoverflow.com/questions/58242892/go-check-if-file-is-a-text-file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Read the first 5 lines
	scanner := bufio.NewScanner(file)
	text := ""
	lineCount := 0
	for scanner.Scan() {
		if lineCount < 5 {
			t := scanner.Text()
			text += t
			lineCount++
		}
	}

	return utf8.ValidString(string(text))
}

func getFileSize(fileName string) int64 {
	stat, err := os.Lstat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return stat.Size()
}

func formatFileSize(size int64) string {
	var fileSizeStr string
	const (
		_          = iota
		kb float64 = 1 << (10 * iota)
		mb
		gb
	)

	switch {
	case size < int64(kb):
		fileSizeStr = fmt.Sprintf("%d bytes", size)
	case size < int64(mb):
		fileSizeStr = fmt.Sprintf("%.2f KB", float64(size)/kb)
	case size < int64(gb):
		fileSizeStr = fmt.Sprintf("%.2f MB", float64(size)/mb)
	default:
		fileSizeStr = fmt.Sprintf("%.2f GB", float64(size)/gb)
	}

	return fileSizeStr
}

func createFileInfoArray(li []string) []FileInfo {
	arr := make([]FileInfo, len(li))

	for i := 0; i < len(li); i++ {
		if isFileReadable(li[i]) == false {
			continue
		}

		arr[i].name = li[i]
		// Attempt to read file and just continue if failed

		arr[i].isSymlink = isSymlink(arr[i].name)

		if arr[i].isSymlink {
			realPath, err := filepath.EvalSymlinks(arr[i].name)
			if err != nil {
				log.Fatal(err)
			}
			arr[i].isDirectory = isDirectory(realPath)
		} else {
			arr[i].isDirectory = isDirectory(arr[i].name)
		}

		if arr[i].isDirectory {
			arr[i].isEmpty = isDirectoryEmpty(arr[i].name)
		} else {
			if !arr[i].isSymlink {
				arr[i].size = getFileSize(arr[i].name)
				arr[i].extension = getFileExtension(arr[i].name)
				arr[i].isEmpty = isFileEmpty(arr[i].name)
				arr[i].isExecutable = isFileExecutable(arr[i].name)
				arr[i].isTextFile = isTextFile(arr[i].name)
			}
		}
	}

	return arr
}
