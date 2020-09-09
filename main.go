package main

import (
	"flag"
	"log"
	"path/filepath"
)

var (
	targetDir  string
	passwd     string
	outputName string
)

func main() {
	flag.StringVar(&targetDir, "d", "", "target directory")
	flag.StringVar(&outputName, "o", "archive.zip", "output file name")
	flag.StringVar(&passwd, "p", "", "password")
	flag.Parse()
	if targetDir == "" {
		log.Fatal("Error obtaining target directory: specify directory with -d")
	}

	var err error
	targetDir, err = filepath.Abs(targetDir)
	if err != nil {
		log.Fatal("Invalid filepath")
	}

	if buildArchive() != nil {
		log.Fatal("Error building an archive")
	}
}
