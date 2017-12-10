package servers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/domain/entities"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/usecases/ports"
)

type DownloadFileHandler struct {
	BaseFileHandler
}

func newDownloadFileHandler(config *config.Config, inputPort ports.ServerInputPort) http.Handler {
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	return withLogging(config, logger, &DownloadFileHandler{
		newBaseFileHandler(logger, inputPort),
	})
}

func (h *DownloadFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p := r.URL.Path

	// allowed format: /v1/files/token.hash
	sp := strings.Split(p, "/")
	if len(sp) != 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.inputPort.DownloadFile(&ports.DownloadFileParams{
		Token: entities.FileID(sp[3]),
	})

	if err != nil {
		h.logger.Printf("[ERR] %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(res.Content)))
	if _, err := w.Write(res.Content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}
