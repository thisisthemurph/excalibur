package handler

import (
	"errors"
	"excalibur/internal/handler/dto"
	"excalibur/internal/handler/errorhandler"
	"excalibur/internal/handler/response"
	"excalibur/internal/services"
	"fmt"
	"log"
	"net/http"
)

// FileHandler handler for files
type FileHandler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	GetFileStatus(w http.ResponseWriter, r *http.Request)
}

type fileHandler struct {
	eh      errorhandler.ErrorHandler
	log     log.Logger
	service services.FileService
}

// NewFileHandler implements a new file handler
func NewFileHandler(s services.FileService, logger log.Logger) FileHandler {
	eh := errorhandler.New()
	return &fileHandler{
		eh:      eh,
		log:     logger,
		service: s,
	}
}

func (f *fileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("fileUpload")

	if err != nil {
		fmt.Println("Issue getting file from request")
		response.ReturnError(w, errors.New("Issue uploading file to server"), http.StatusInternalServerError)
		return
	}

	err = f.service.UploadFile(file, header)
	if err != nil {
		response.ReturnError(w, err, http.StatusInternalServerError)
		return
	}

	response.Respond(w, dto.FileUploadStatusDTO{Status: "Uploading"}, http.StatusAccepted)
}

func (f *fileHandler) GetFileStatus(w http.ResponseWriter, r *http.Request) {
	uploadStatus := dto.FileUploadStatusDTO{
		Status: "Uploading",
	}

	fmt.Println(uploadStatus)

	response.Respond(w, uploadStatus, http.StatusOK)
}
