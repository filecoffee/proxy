package modules

import (
	"log"
	"net/http"
)

const Base string = "https://file.coffee/u/"

func GetUpload(upload string) *http.Response {
	res, err := http.Get(Base + upload)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
