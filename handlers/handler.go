package handlers

import (
	"net/http"

	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/log"
)

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

func NewFileHandler(config *config.Config) http.Handler {
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
