package presenters

import (
	"io"
	"io/ioutil"

	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/domain/entities"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/usecases/ports"
)

type HTTPPresenter struct {
	logger *log.Logger
	config *config.Config
}

func NewHTTPPresenter(logger *log.Logger, config *config.Config) *HTTPPresenter {
	return &HTTPPresenter{
		logger: logger,
		config: config,
	}
}

func (p *HTTPPresenter) UploadFile(token entities.FileID) (*ports.UploadFileResponse, error) {
	return &ports.UploadFileResponse{
		Token: string(token),
	}, nil
}

func (p *HTTPPresenter) DownloadFile(content io.Reader) (*ports.DownloadFileResponse, error) {
	res := &ports.DownloadFileResponse{}
	var err error
	res.Content, err = ioutil.ReadAll(content)
	return res, err
}
