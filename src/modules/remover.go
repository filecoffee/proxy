package modules

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Setup() {
	if os.Getenv("CACHE") == "false" {
		return
	}
	c := cron.New()
	_, _ = c.AddFunc("* * * * *", func() { checkFiles() })
	c.Start()
}

func checkFiles() {
	log.Print("[PROXY] Checking all uploads")
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
	days, _ := strconv.ParseInt(os.Getenv("CACHE_DAYS"), 10, 64)
	if currentTime-fileTime > days*int64(24*60*60*1000) {
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
