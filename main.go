package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// By default don't include hidden files
	showHidden := false

	if len(os.Args) > 1 {
		if os.Args[1] == "-a" {
			showHidden = true
		} else {
			log.Fatal("ERROR: Unrecognized argument")
		}
	}

	cwdContent, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	fileList := make([]string, 0)
	for _, content := range cwdContent {
		if !showHidden && strings.HasPrefix(content.Name(), ".") {
			continue
		}
		fileList = append(fileList, content.Name())
	}

	fileInfoArray := createFileInfoArray(fileList)
	sort.Sort(ByDirectoryAndName(fileInfoArray))

	printIconAndFilenames(fileInfoArray)
}
