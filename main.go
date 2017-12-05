package main

import (
	"github.com/k0kubun/pp"
	"github.com/ktr0731/cris/adapters/servers"
	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/log"
)

func main() {
	config := config.Get()
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	pp.Println(config)

	if err := servers.NewHTTPServer(logger, config).Listen(); err != nil {
		panic(err)
	}
}
