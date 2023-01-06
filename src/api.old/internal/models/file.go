package models

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FileUploadStatus details the status of a file upload
type FileUploadStatus string

const (
	// FileStatusUploading - a file that is currently being uploaded
	FileStatusUploading FileUploadStatus = "Uploading"
	// FileStatusUploadComplete - a file that has successfully uploaded
	FileStatusUploadComplete = "UploadComplete"
	// FileStatusUploadFailed - the upload has failed
	FileStatusUploadFailed = "UploadFailed"
	// FileStatusUnknown - the status of the file is not known
	FileStatusUnknown = "Unknown"
)

// FileMetadata represents a file that has been uploaded
type FileMetadata struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	StoredName string             `bson:"storedName" json:"storedName"`
	Status     string             `bson:"status" json:"status"`
}

func (s FileUploadStatus) String() string {
	switch s {
	case FileStatusUploading:
		return "Uploading"
	case FileStatusUploadComplete:
		return "UploadComplete"
	}

	return "Unknown"
}

// GetFileUploadStatus return a the status enum of the string or `Unknown`
func GetFileUploadStatus(s string) FileUploadStatus {
	switch s {
	case "Uploading":
		return FileStatusUploading
	case "UploadComplete":
		return FileStatusUploadComplete
	}

	return FileStatusUnknown
}

// MakeNewFile represents an object of a new file to be uploaded
type MakeNewFile struct {
	File     multipart.File
	Filename string
}
