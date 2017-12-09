package servers

import (
	"fmt"
	"net/http"

	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/usecases"
	"github.com/ktr0731/cris/server/usecases/ports"
	"github.com/rs/cors"
)

type Server struct {
	logger *log.Logger
	config *config.Config

	mux *http.ServeMux
}

func NewHTTPServer(logger *log.Logger, config *config.Config, inputPort ports.ServerInputPort) *Server {
	s := Server{
		logger: logger,
		config: config,
		mux:    http.NewServeMux(),
	}
	s.mux.Handle(s.getPrefix()+"/files", newUploadFileHandler(config, inputPort))
	s.mux.Handle(s.getPrefix()+"/files/", newDownloadFileHandler(config, inputPort))
	return &s
}

func (s *Server) Listen() error {
	s.logger.Printf("Server listen in %s%s", s.getPrefix(), s.getAddr())
	return http.ListenAndServe(s.getAddr(), cors.Default().Handler(s.mux))
}

func (s *Server) getPrefix() string {
	return fmt.Sprintf("/%s", s.config.Meta.Version)
}

func (s *Server) getAddr() string {
	return fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)
}

type BaseFileHandler struct {
	logger *log.Logger

	inputPort ports.ServerInputPort
}

func newBaseFileHandler(logger *log.Logger, inputPort ports.ServerInputPort) BaseFileHandler {
	return BaseFileHandler{
		logger:    logger,
		inputPort: inputPort,
	}
}

func handleError(w http.ResponseWriter, err error) {
	if _, ok := err.(usecases.ClientError); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintln(w, err)
}
