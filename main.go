package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/handlers"
)

func main() {
	config := config.Get()

	pp.Println(config)

	prefix := fmt.Sprintf("/%s", config.Meta.Version)

	mux := http.NewServeMux()
	mux.Handle(prefix+"/files", handlers.NewFileHandler(config))

	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("Server listen in %s%s", addr, prefix)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}
