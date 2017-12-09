package servers

import (
	"net/http"

	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/log"
)

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
