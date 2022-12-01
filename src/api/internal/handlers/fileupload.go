package handlers

import (
	"excalibur/internal/upload"
	"fmt"
	"io"
	"net/http"
)

// FileUploadHandler a handler for uploading a file
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("fileUpload")

	if err != nil {
		fmt.Println("Issue getting file from request")
		panic(err)
	}

	result := upload.UploadFile(file, header)
	io.WriteString(w, result)
}
