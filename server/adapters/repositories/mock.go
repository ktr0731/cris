package repositories

import (
	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/domain/entities"
	"github.com/ktr0731/cris/domain/repositories"
	"github.com/ktr0731/cris/log"
	"golang.org/x/sync/syncmap"
)

type MockFileRepositoryAdapter struct {
	logger *log.Logger
	config *config.Config

	storage syncmap.Map
}

func NewMockFileRepository(logger *log.Logger, config *config.Config) *MockFileRepositoryAdapter {
	return &MockFileRepositoryAdapter{
		logger: logger,
		config: config,
	}
}

func (r *MockFileRepositoryAdapter) Store(e *entities.File) (entities.FileID, error) {
	r.storage.Store(e.ID, e)
	return e.ID, nil
}

func (r *MockFileRepositoryAdapter) Find(id entities.FileID) (*entities.File, error) {
	v, ok := r.storage.Load(id)
	if !ok {
		return nil, repositories.ErrNotFound
	}
	e, ok := v.(*entities.File)
	if !ok {
		return nil, repositories.ErrNotFound
	}
	return e, nil
}

func (r *MockFileRepositoryAdapter) Remove(id entities.FileID) (*entities.File, error) {
	e, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.storage.Delete(id)
	return e, nil
}
