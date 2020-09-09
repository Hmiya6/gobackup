package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// filepath.WalkFunc wrapper
func visitDirs(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		// skip method is NOT COMPLETED. Need more work
		// skip files as "directory".
		if info.IsDir() {
			return nil
		}

		// skip files in ".git"
		if strings.Contains(path, ".git") {
			return nil
		}

		// skip executable files
		if filepath.Ext(path) == "" {
			return nil
		}

		// add file to files list
		*files = append(*files, path)
		return err
	}
}

// retreiveFiles() obtain file list to archive.
func retreiveFiles() ([]string, error) {
	files := []string{}

	// walk in targetDir recursively
	err := filepath.Walk(targetDir, visitDirs(&files))
	if err != nil {
		newErrStr := fmt.Sprintf("Error while visiting directories: %s", err.Error())
		err = errors.New(newErrStr)
		return nil, err
	}

	return files, nil
}

// addFiles() adds files from file list into a zip file
func addFiles(zipWriter *zip.Writer, files []string) error {

	// do the following against all file
	for _, file := range files {

		// read file data
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		// remove superfluos directory structure
		relativePath := strings.TrimPrefix(file, targetDir)

		// add file in zip file
		f, err := zipWriter.Create(relativePath)
		if err != nil {
			return err
		}

		// write content into file in .zip
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}

// compressFiles() makes zip file from file list
func compressFiles(files []string) error {
	// set output file
	outFile, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// set zip writer
	zipWriter := zip.NewWriter(outFile)

	// add files into zip
	err = addFiles(zipWriter, files)
	if err != nil {
		return err
	}

	defer zipWriter.Close()
	return nil
}

// buildArchive() wraps all functions to generate archive.
func buildArchive() error {
	files, err := retreiveFiles()
	if err != nil {
		return err
	}
	if compressFiles(files) != nil {
		return err
	}
	return nil
}
