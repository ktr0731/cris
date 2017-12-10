package servers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/log"
	"github.com/ktr0731/cris/server/usecases/ports"
)

type UploadFileHandler struct {
	BaseFileHandler
}

func newUploadFileHandler(config *config.Config, inputPort ports.ServerInputPort) http.Handler {
	logger, err := log.NewLogger(config)
	if err != nil {
		panic(err)
	}
	return withLogging(config, logger, &UploadFileHandler{
		newBaseFileHandler(logger, inputPort),
	})
}

func (h *UploadFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	params := struct {
		Content   string `json:"content"`
		Signature string `json:"signature"`
		Pubkey    string `json:"pubkey"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		h.logger.Printf("[ERR] %s", err)
		handleError(w, err)
		return
	}

	content, pubkey, signature := make([]byte, base64.StdEncoding.DecodedLen(len(params.Content))), make([]byte, base64.StdEncoding.DecodedLen(len(params.Pubkey))), make([]byte, base64.StdEncoding.DecodedLen(len(params.Signature)))
	var err error
	var n int
	n, err = base64.StdEncoding.Decode(content, []byte(params.Content))
	if err != nil {
		h.logger.Printf("[ERR] %s", err)
		handleError(w, err)
		return
	}
	content = content[:n]

	n, err = base64.StdEncoding.Decode(pubkey, []byte(params.Pubkey))
	if err != nil {
		h.logger.Printf("[ERR] %s", err)
		handleError(w, err)
		return
	}
	pubkey = pubkey[:n]

	n, err = base64.StdEncoding.Decode(signature, []byte(params.Signature))
	if err != nil {
		h.logger.Printf("[ERR] %s", err)
		handleError(w, err)
		return
	}
	signature = signature[:n]

	fmt.Printf("%x, %x", pubkey, sha256.Sum256(content))

	res, err := h.inputPort.UploadFile(&ports.UploadFileParams{
		Content:   bytes.NewBuffer(content),
		Signature: signature,
		PubKey:    pubkey,
	})

	if err != nil {
		h.logger.Printf("[ERR] %s", err)
		handleError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
}
