package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	name         string
	isSymlink    bool
	isDirectory  bool
	isEmpty      bool
	size         int64
	extension    string
	isExecutable bool
}

type ByDirectoryAndName []FileInfo

func (a ByDirectoryAndName) Len() int      { return len(a) }
func (a ByDirectoryAndName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDirectoryAndName) Less(i, j int) bool {
	if a[i].isDirectory && !a[j].isDirectory {
		return true
	} else if !a[i].isDirectory && a[j].isDirectory {
		return false
	} else {
		return a[i].name < a[j].name
	}
}

func isTextFile(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Read up to 512 bytes (DetectContentType uses at most 512 bytes)
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return false, err
	}

	mimeType := http.DetectContentType(buf[:n])
	// MIME type for plain text is usually "text/plain; charset=utf-8"
	return mimeType == "text/plain; charset=utf-8" || mimeType[:5] == "text/", nil
}

func getFileInfo(path string) FileInfo {
	var info FileInfo
	info.name = path

	fi, err := os.Lstat(path)
	if err != nil {
		// unreadable file, just return name
		return info
	}

	mode := fi.Mode()
	info.size = fi.Size()
	info.isExecutable = (mode & 0111) != 0
	info.isDirectory = mode.IsDir()
	info.extension = strings.TrimPrefix(filepath.Ext(path), ".")

	// symlink?
	if mode&os.ModeSymlink != 0 {
		info.isSymlink = true
		if realPath, err := filepath.EvalSymlinks(path); err == nil {
			if st, err := os.Stat(realPath); err == nil {
				info.isDirectory = st.IsDir()
				info.size = st.Size()
			}
		}
	}

	// expensive checks → only if really needed
	if info.isDirectory {
		// check emptiness by reading just 1 entry
		if f, err := os.Open(path); err == nil {
			defer f.Close()
			if _, err := f.Readdirnames(1); err == io.EOF {
				info.isEmpty = true
			}
		}
	} else {
		if info.size == 0 {
			info.isEmpty = true
		}
	}

	return info
}

func createFileInfoArray(files []string) []FileInfo {
	arr := make([]FileInfo, len(files))
	for i, f := range files {
		arr[i] = getFileInfo(f)
	}
	return arr
}
