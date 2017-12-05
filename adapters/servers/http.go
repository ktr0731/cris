package servers

import (
	"fmt"
	"net/http"

	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/log"
)

type Server struct {
	logger *log.Logger
	config *config.Config

	mux *http.ServeMux
}

func NewHTTPServer(logger *log.Logger, config *config.Config) *Server {
	s := Server{
		logger: logger,
		config: config,
		mux:    http.NewServeMux(),
	}
	s.mux.Handle(s.getPrefix()+"/files", newFileHandler(config))
	return &s
}

func (s *Server) Listen() error {
	s.logger.Printf("Server listen in %s%s", s.getPrefix(), s.getAddr())
	return http.ListenAndServe(s.getAddr(), s.mux)
}

func (s *Server) getPrefix() string {
	return fmt.Sprintf("/%s", s.config.Meta.Version)
}

func (s *Server) getAddr() string {
	return fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)
}

type FileHandler struct {
	logger   *log.Logger
	upload   http.Handler
	download http.Handler
}

func (h *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.download.ServeHTTP(w, r)

	case http.MethodPost:
		h.upload.ServeHTTP(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func newFileHandler(config *config.Config) http.Handler {
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	return withLogging(config, logger, &FileHandler{
		logger:   logger,
		upload:   &uploadHandler{logger},
		download: &downloadHandler{logger},
	})
}

type uploadHandler struct {
	logger *log.Logger
}

func (h *uploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("upload")
}

type downloadHandler struct {
	logger *log.Logger
}

func (h *downloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("download")
}

type logHandler struct {
	logger  *log.Logger
	handler http.Handler
}

func (h *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("%s %s", r.Method, r.URL.Path)
	h.handler.ServeHTTP(w, r)
}

func withLogging(config *config.Config, logger *log.Logger, h http.Handler) http.Handler {
	return &logHandler{
		logger:  logger,
		handler: h,
	}
}
