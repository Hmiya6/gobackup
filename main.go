package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

var (
	targetDir  string
	outputName string
)

func main() {
	flag.StringVar(&targetDir, "d", "", "target directory")
	now := time.Now()
	defaultOut := fmt.Sprintf("archive-%v-%v-%v.zip", now.Year(), int(now.Month()), now.Day())
	flag.StringVar(&outputName, "o", defaultOut, "output file name")
	flag.StringVar(&credentialJSONFile, "c", "credentials.json", "Google OAuth2 credentials (.json file)")
	flag.Parse()
	if targetDir == "" {
		log.Fatal("Error obtaining target directory: specify directory with -d")
	}

	var err error
	targetDir, err = filepath.Abs(targetDir)
	if err != nil {
		log.Fatal("Invalid filepath")
	}
	fmt.Printf("archive %s\n", targetDir)

	err = buildArchive()
	if err != nil {
		log.Fatal("Error building an archive:", err)
	}

	err = uploadFile(outputName, resumableUpload)
	if err != nil {
		log.Fatal("Error uploading a file:", err)
	}
}
