package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	showHidden := flag.Bool("a", false, "Show hidden files")

	flag.Parse()

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	files := []string{}
	for _, v := range entries {
		if *showHidden || !strings.HasPrefix(v.Name(), ".") {
			files = append(files, v.Name())
		}
	}

	fileInfo := createFileInfoArray(files)
	sort.Sort(ByDirectoryAndName(fileInfo))

	for _, v := range fileInfo {
		printIconFilename(v)
	}
}
