package modules

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func CacheSetup() {
	if os.Getenv("CACHE") == "false" {
		return
	}
	pwd, _ := os.Getwd()
	if _, err := os.Stat(filepath.Join(pwd, "../cache/")); os.IsNotExist(err) {
		_ = os.Mkdir(filepath.Join(pwd, "../cache/"), os.ModePerm)
	}
}

func Cache(upload string, request *http.Response) (err error) {
	if os.Getenv("CACHE") == "false" {
		return
	}
	md5Hash := stringToMd5(upload)

	pwd, _ := os.Getwd()
	out, err := os.Create(filepath.Join(pwd, "../cache/"+md5Hash))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, request.Body)
	if err != nil {
		return err
	}
	return nil
}

func GetFromCache(upload string) (result []byte, err error) {
	if os.Getenv("CACHE") == "false" {
		return nil, nil
	}
	md5Hash := stringToMd5(upload)
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(filepath.Join(pwd, "../cache/"+md5Hash))
	if err != nil {
		return nil, err
	}
	return data, nil

}

func CheckIfCached(upload string) (result bool) {
	md5Hash := stringToMd5(upload)
	pwd, _ := os.Getwd()
	info, err := os.Stat(filepath.Join(pwd, "../cache/"+md5Hash))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()

}

func stringToMd5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
