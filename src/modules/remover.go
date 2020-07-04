package modules

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func Setup() {
	checkFiles()
}

func checkFiles() {
	files := getFiles()
	for _, file := range files {
		checkFile(file)

	}
}

func checkFile(upload string) {
	file, err := os.Stat(upload)
	if err != nil {
		log.Print(err)
	}

	modifiedTime := file.ModTime()

	fileTime := modifiedTime.UnixNano() / int64(time.Millisecond)
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	if currentTime-fileTime > 7*86400000 {
		_ = os.Remove(upload)
	}
}

func getFiles() []string {
	pwd, _ := os.Getwd()
	var base string = filepath.Join(pwd, "../cache/")
	var files []string

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
