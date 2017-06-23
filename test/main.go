package main

import (
	"log"

	"github.com/slashquery/download"
)

const URL = "https://codeload.github.com/slashquery/slashquery/zip/master"

func main() {
	err := download.Download(URL, "/tmp/sq.zip")
	if err != nil {
		log.Fatal(err)
	}
	err = download.Unzip("/tmp/sq.zip", "/tmp/sq")
}
