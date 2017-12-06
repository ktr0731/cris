package usecases

import (
	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/domain/entities"
	"github.com/ktr0731/cris/domain/repositories"
	"github.com/ktr0731/cris/log"
	"github.com/ktr0731/cris/usecases/ports"
	"github.com/ktr0731/cris/utils"
)

type container struct {
	config *config.Config
	logger *log.Logger
}

// uploadFile uploads received file to an object storage
func (c *container) uploadFile(
	params *ports.UploadFileParams,
	outputPort ports.ServerOutputPort,
	storagePort ports.StoragePort,
	fileRepository repositories.FileRepository,
) (*ports.UploadFileResponse, error) {
	url, err := storagePort.Upload(utils.NewUUID(), params.Content)
	if err != nil {
		return nil, err
	}

	file := entities.NewFile(url)
	_, err = fileRepository.Store(file)
	if err != nil {
		return nil, err
	}

	return outputPort.UploadFile(file.ID)
}

func newUsecaseContainer(logger *log.Logger, config *config.Config) *container {
	return &container{
		logger: logger,
		config: config,
	}
}
