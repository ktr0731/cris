package entities

import "github.com/ktr0731/cris/utils"

type FileID string

type File struct {
	ID  FileID
	URL string
}

func NewFile(url string) *File {
	return &File{
		ID:  FileID(utils.NewUUID()),
		URL: url,
	}
}
