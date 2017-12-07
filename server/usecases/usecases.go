package usecases

import (
	"io/ioutil"

	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/domain/entities"
	"github.com/ktr0731/cris/server/domain/repositories"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/usecases/ports"
	"github.com/ktr0731/cris/server/utils"
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

func (c *container) downloadFile(
	params *ports.DownloadFileParams,
	outputPort ports.ServerOutputPort,
	storagePort ports.StoragePort,
	blockchainPort ports.BlockchainPort,
	cryptoPort ports.CryptoPort,
	fileRepository repositories.FileRepository,
) (*ports.DownloadFileResponse, error) {
	tx, err := blockchainPort.FindTxByHash(params.TxHash)
	if err != nil {
		return nil, err
	}

	file, err := fileRepository.Find(params.Token)
	if err != nil {
		return nil, err
	}

	res, err := storagePort.Download(file.URL)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	b, err := ioutil.ReadAll(res)
	if err != nil {
		return nil, err
	}

	if cryptoPort.HashDigest(b) != tx.HashedData {
		return nil, ErrTemperingDetected
	}

	return &ports.DownloadFileResponse{
		Content: b,
	}, nil
}

func newUsecaseContainer(logger *log.Logger, config *config.Config) *container {
	return &container{
		logger: logger,
		config: config,
	}
}
