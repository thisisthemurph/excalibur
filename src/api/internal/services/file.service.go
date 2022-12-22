package services

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

// FileService functionality for handling files
type FileService interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader) error
}

type fileService struct {
	log log.Logger
}

func newFileService(l log.Logger) FileService {
	return &fileService{
		log: l,
	}
}

func getFileExtension(fn string) (string, bool) {
	parts := strings.Split(fn, ".")
	if len(parts) < 2 {
		return "", false
	}

	return parts[len(parts)-1], true
}

func getUploadedFileName(fn string) (string, error) {
	time := time.Now().UTC().Format(time.RFC3339)
	ext, ok := getFileExtension(fn)
	if !ok {
		return "", errors.New("The file must have a file extension")
	}

	name := fmt.Sprintf("UploadedFile_%v.%v", time, ext)
	return name, nil
}

func (fs *fileService) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) error {
	fmt.Println("Uploading a file")

	fileName, err := getUploadedFileName(fileHeader.Filename)
	if err != nil {
		return err
	}

	if !strings.HasSuffix(strings.ToLower(fileName), "csv") {
		return errors.New("Only .csv files are accepted")
	}

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Issue opening file")
		fmt.Println(err)
		return errors.New("The uploaded file could not be opened")
	}

	defer f.Close()
	_, _ = io.Copy(f, file)

	return nil
}
