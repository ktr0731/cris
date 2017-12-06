package main

import (
	"github.com/k0kubun/pp"
	"github.com/ktr0731/cris/adapters/blockchains"
	"github.com/ktr0731/cris/adapters/crypto"
	"github.com/ktr0731/cris/adapters/repositories"
	"github.com/ktr0731/cris/adapters/servers"
	"github.com/ktr0731/cris/adapters/storages"
	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/log"
	"github.com/ktr0731/cris/presenters"
	"github.com/ktr0731/cris/usecases"
)

func main() {
	config := config.Get()
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	pp.Println(config)

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