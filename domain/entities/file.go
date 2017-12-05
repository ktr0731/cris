package entities

import uuid "github.com/satori/go.uuid"

type FileID string

type File struct {
	ID  FileID
	URL string
}

func NewFile(url string) *File {
	return &File{
		ID:  FileID(uuid.NewV4().String()),
		URL: url,
	}
}
