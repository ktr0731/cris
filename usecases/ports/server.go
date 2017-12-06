package ports

import (
	"io"

	"github.com/ktr0731/cris/domain/entities"
)

type ServerInputPort interface {
	UploadFile(*UploadFileParams) (*UploadFileResponse, error)
	DownloadFile(*DownloadFileParams) (*DownloadFileResponse, error)
}

type UploadFileParams struct {
	Content io.Reader
}

type DownloadFileParams struct {
	Token  entities.FileID
	TxHash string
}

type ServerOutputPort interface {
	UploadFile(token entities.FileID) (*UploadFileResponse, error)
	DownloadFile(content io.Reader) (*DownloadFileResponse, error)
}

type UploadFileResponse struct {
	Token string `json:"token"`
}

type DownloadFileResponse struct {
	Content []byte `json:"content"`
}
