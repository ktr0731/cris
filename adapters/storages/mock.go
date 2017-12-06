package storages

import (
	"errors"
	"io"

	"github.com/ktr0731/cris/domain/repositories"

	"golang.org/x/sync/syncmap"
)

type MockStorageAdapter struct {
	storage syncmap.Map
}

func (s *MockStorageAdapter) Upload(name string, content io.Reader) (string, error) {
	s.storage.Store(name, content)
	return "", nil
}

func (s *MockStorageAdapter) Download(name string) (io.Reader, error) {
	v, ok := s.storage.Load(name)
	if !ok {
		return nil, repositories.ErrNotFound
	}
	r, ok := v.(io.Reader)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return r, nil
}

func NewMockStorage() *MockStorageAdapter {
	return &MockStorageAdapter{}
}
