package main

import (
	"github.com/ktr0731/cris/server/adapters/blockchains"
	"github.com/ktr0731/cris/server/adapters/crypto"
	"github.com/ktr0731/cris/server/adapters/repositories"
	"github.com/ktr0731/cris/server/adapters/servers"
	"github.com/ktr0731/cris/server/adapters/storages"
	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/presenters"
	"github.com/ktr0731/cris/server/usecases"
)

func main() {
	config := config.Get()
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	interactor := usecases.NewInteractor(
		logger,
		config,
		presenters.NewHTTPPresenter(logger, config),
		storages.NewMockStorage(),
		blockchains.NewMockBlockchain(),
		crypto.NewCryptoAdapter(),
		repositories.NewMockFileRepository(logger, config),
	)

	if err := servers.NewHTTPServer(logger, config, interactor).Listen(); err != nil {
		panic(err)
	}
}
