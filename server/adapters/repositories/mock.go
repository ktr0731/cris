package repositories

import (
	"github.com/k0kubun/pp"
	"github.com/ktr0731/cris/server/config"
	"github.com/ktr0731/cris/server/domain/entities"
	"github.com/ktr0731/cris/server/domain/repositories"
	"github.com/ktr0731/cris/server/log"
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
	r.logger.Printf("[MockFileRepo] stored an entity: %s", e.ID)
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
	pp.Println(e)
	return e, nil
}

func (r *MockFileRepositoryAdapter) Remove(id entities.FileID) (*entities.File, error) {
	e, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.storage.Delete(id)
	r.logger.Printf("[MockFileRepo] remove an entity: %s", e.ID)
	return e, nil
}
