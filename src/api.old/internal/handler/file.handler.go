package handler

import (
	"errors"
	"excalibur/internal/handler/dto"
	"excalibur/internal/handler/errorhandler"
	"excalibur/internal/handler/request"
	"excalibur/internal/handler/response"
	"excalibur/internal/models"
	"excalibur/internal/services"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// FileHandler handler for files
type FileHandler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	GetFileStatus(w http.ResponseWriter, r *http.Request)
}

type fileHandler struct {
	eh  errorhandler.ErrorHandler
	log log.Logger
	fs  services.FileService
	dts services.DataTemplateService
}

// NewFileHandler implements a new file handler
func NewFileHandler(fs services.FileService, dts services.DataTemplateService, logger log.Logger) FileHandler {
	eh := errorhandler.New()
	return &fileHandler{
		eh:  eh,
		log: logger,
		fs:  fs,
		dts: dts,
	}
}

func getFileExtension(fn string) (string, bool) {
	parts := strings.Split(fn, ".")
	if len(parts) < 2 {
		return "", false
	}

	return parts[len(parts)-1], true
}

func makeUploadFilename(filename string, ext string) string {
	time := time.Now().UTC().Format(time.RFC3339)
	return fmt.Sprintf("UploadedFile_%v.%v", time, ext)
}

func (f *fileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("fileUpload")
	dataTemplateID, idOk := request.GetParamFomRequest(r, "id")

	if err != nil {
		fmt.Println("Issue getting file from request")
		response.ReturnError(w, errors.New("issue uploading file to server"), http.StatusInternalServerError)
		return
	}

	if !idOk {
		response.ReturnError(w, errors.New("could not determine the DataTemplate ID"), http.StatusBadRequest)
		return
	}

	// Determine the appropriate file name for the upload

	ext, hasExt := getFileExtension(header.Filename)

	if !hasExt {
		response.ReturnError(w, errors.New("the file must have an extension"), http.StatusInternalServerError)
		return
	}

	if ext != "csv" {
		response.ReturnError(w, errors.New("the file must have a .csv extension"), http.StatusInternalServerError)
		return
	}

	filename := makeUploadFilename(header.Filename, ext)

	// Start the file upload

	fileChan := make(chan error)
	newFile := models.MakeNewFile{
		File:     file,
		Filename: filename,
	}

	go f.fs.UploadFile(fileChan, newFile)

	// Upload the file metadata to the database

	metadata := models.FileMetadata{
		Name:       header.Filename,
		StoredName: filename,
		Status:     fmt.Sprint(models.FileStatusUploading),
	}

	insertedMetadata, err := f.dts.AddFileMetadata(dataTemplateID, metadata)
	if err != nil {
		response.ReturnError(w, errors.New("failed to add file metadata to database"), http.StatusInternalServerError)
		return
	}

	// Return status stating that the file is being uploaded

	status := dto.FileUploadStatusDTO{
		FileID: insertedMetadata.ID.String(),
		Status: models.GetFileUploadStatus(insertedMetadata.Status),
	}

	response.Respond(w, status, http.StatusAccepted)

	// Handle any file upload failures

	fileUploadErr := <-fileChan
	if fileUploadErr != nil {
		f.log.Println("There was an error uploading the file")
		f.dts.UpdateFileStatus(dataTemplateID, status.FileID, models.FileStatusUploadFailed) // TODO
	} else {
		f.log.Println("The file has been uploaded")
		f.dts.UpdateFileStatus(dataTemplateID, status.FileID, models.FileStatusUploadComplete)
	}
}

func (f *fileHandler) GetFileStatus(w http.ResponseWriter, r *http.Request) {
	uploadStatus := dto.FileUploadStatusDTO{
		Status: models.FileStatusUploading,
	}

	fmt.Println(uploadStatus)

	response.Respond(w, uploadStatus, http.StatusOK)
}
