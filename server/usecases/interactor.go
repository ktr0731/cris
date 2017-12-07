package usecases

import (
	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/domain/repositories"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/usecases/ports"
)

// interactor はすべての依存を解決する役割を持つ
type Interactor struct {
	container *container

	outputPort     ports.ServerOutputPort
	storagePort    ports.StoragePort
	cryptoPort     ports.CryptoPort
	fileRepository repositories.FileRepository
}

func (i *Interactor) UploadFile(params *ports.UploadFileParams) (*ports.UploadFileResponse, error) {
	return i.container.uploadFile(params, i.outputPort, i.storagePort, i.fileRepository)
}

// DownloadFile downloads a file specified by the token
func (i *Interactor) DownloadFile(params *ports.DownloadFileParams) (*ports.DownloadFileResponse, error) {
	return i.outputPort.DownloadFile(nil)
}

func NewInteractor(
	logger *log.Logger,
	config *config.Config,
	outputPort ports.ServerOutputPort,
	storagePort ports.StoragePort,
	blockchainPort ports.BlockchainPort,
	cryptoPort ports.CryptoPort,
	fileRepository repositories.FileRepository,
) *Interactor {
	return &Interactor{
		container: newUsecaseContainer(logger, config),

		outputPort:     outputPort,
		storagePort:    storagePort,
		cryptoPort:     cryptoPort,
		fileRepository: fileRepository,
	}
}
