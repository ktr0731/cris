package servers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/log"
	"github.com/ktr0731/cris/usecases"
	"github.com/ktr0731/cris/usecases/ports"
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
	s.mux.Handle(s.getPrefix()+"/files", newFileHandler(config, inputPort))
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
	logger *log.Logger

	inputPort ports.ServerInputPort
}

func newFileHandler(config *config.Config, inputPort ports.ServerInputPort) http.Handler {
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	return withLogging(config, logger, &FileHandler{
		logger:    logger,
		inputPort: inputPort,
	})
}

func (h *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var err error
	switch r.Method {
	case http.MethodPost:
		res, err = h.uploadFile(w, r)
	case http.MethodGet:
		h.downloadFile(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		handleError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}

func (h *FileHandler) uploadFile(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	defer r.Body.Close()

	return h.inputPort.UploadFile(&ports.UploadFileParams{
		Content: r.Body,
	})
}

func (h *FileHandler) downloadFile(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path

	// allowed format: /v1/files/token.hash.privkey
	sp := strings.Split(p, "/")
	if len(sp) != 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vals := strings.Split(sp[2], ".")
	if len(vals) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	privKey, err := base64.StdEncoding.DecodeString(vals[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.inputPort.DownloadFile(&ports.DownloadFileParams{
		Token:   vals[0],
		TxHash:  vals[1],
		PrivKey: privKey,
	})

	return
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

func handleError(w http.ResponseWriter, err error) {
	pp.Println(err)
	if _, ok := err.(usecases.ClientError); ok {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintln(w, err)
}
