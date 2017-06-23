package main

import (
	"log"

	"github.com/slashquery/download"
)

func main() {
	err := download.Download("https://github.com/slashquery/slashquery/archive/master.zip", "/tmp/sq.zip")
	if err != nil {
		log.Fatal(err)
	}
}
