package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file multipart.File, fileHeader *multipart.FileHeader) string {
	fmt.Println("Uploading a file")

	f, err := os.OpenFile(fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Issue opening file")
		panic(err)
	}

	defer f.Close()
	_, _ = io.Copy(f, file)
	return "File " + fileHeader.Filename + " uploaded successfully"
}