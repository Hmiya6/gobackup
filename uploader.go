package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
)

const (
	// use "siple upload" when
	simpleUpload    = "media"
	multipartUpload = "multipart"
	resumableUpload = "resumable"
)

/*
func createDir(service *drive.Service, name string, parentId string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
	}

	file, err := service.Files.Create(d).Do()
	if err != nil {
		log.Println("Could not create dri: " + err.Error())
	}

	return file, nil
}
*/

func createFile(service *drive.Service, name string, mimeType string,
	content io.Reader, parentId string) (*drive.File, error) {

	f := &drive.File{
		MimeType: mimeType,
		Name:     name,
		Parents:  []string{parentId},
	}

	fmt.Println("Now uploading file... It may take a few minutes")
	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		log.Println("could not create file: " + err.Error())
		return nil, err
	}
	return file, nil
}

func uploadFile(name string, mimeType string) error {
	f, err := os.Open(name)
	if err != nil {
		log.Printf("Unable to open file: %v\n", err)
		return err
	}
	defer f.Close()

	service, err := getService(credentialJSONFile)

	file, err := createFile(service, name, mimeType, f, "root")
	if err != nil {
		log.Printf("Unable to create file: %v\n", err)
	}

	fmt.Printf("File '%s' successfully uploaded in root\n", file.Name)
	return nil
}

/*
func main() {
	err := uploadFile("archive.zip", resumableUpload)
	if err != nil {
		log.Fatal(err)
	}
}
*/
