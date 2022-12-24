package services

import (
	"errors"
	"excalibur/internal/models"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// FileService functionality for handling files
type FileService interface {
	UploadFile(c chan error, upload models.MakeNewFile)
}

type fileService struct {
	log log.Logger
}

func newFileService(l log.Logger) FileService {
	return &fileService{
		log: l,
	}
}

func (fs *fileService) UploadFile(c chan error, upload models.MakeNewFile) {
	fmt.Println("Uploading a file")

	if !strings.HasSuffix(strings.ToLower(upload.Filename), "csv") {
		c <- errors.New("only .csv files are accepted")
		return
	}

	f, err := os.OpenFile(upload.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c <- errors.New("uploaded file could not be opened")
		return
	}

	defer f.Close()
	_, _ = io.Copy(f, upload.File)

	// return nil
	c <- nil
}
