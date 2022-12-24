package dto

import "excalibur/internal/models"

// FileUploadStatusDTO specifies the status of an uploaded file
type FileUploadStatusDTO struct {
	FileID string                  `json:"fileId"`
	Status models.FileUploadStatus `json:"status"`
}
