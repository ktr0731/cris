package main

import (
	"fmt"
	"net/http"

	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/handlers"
)

func main() {
	config := config.Get()

	prefix := fmt.Sprintf("/%s", config.Meta.Version)

	mux := http.NewServeMux()
	mux.Handle(prefix+"/files", handlers.NewFileHandler())

	if err := http.ListenAndServe(config.Server.Host+config.Server.Port, mux); err != nil {
		panic(err)
	}
}
