package handlers

import "net/http"

type FileHandler struct {
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

func NewFileHandler() *FileHandler {
	return &FileHandler{
		upload:   &uploadHandler{},
		download: &downloadHandler{},
	}
}

type uploadHandler struct {
}

func (h *uploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
}

type downloadHandler struct {
}

func (h *downloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
}
